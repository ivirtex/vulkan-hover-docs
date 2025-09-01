# VK\_EXT\_device\_memory\_report(3) Manual Page

## Name

VK\_EXT\_device\_memory\_report - device extension



## [](#_registered_extension_number)Registered Extension Number

285

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [Developer tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Yiwei Zhang [\[GitHub\]zzyiwei](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_device_memory_report%5D%20%40zzyiwei%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_device_memory_report%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-01-06

**IP Status**

No known IP claims.

**Contributors**

- Yiwei Zhang, Google
- Jesse Hall, Google

## [](#_description)Description

This device extension allows registration of device memory event callbacks upon device creation, so that applications or middleware can obtain detailed information about memory usage and how memory is associated with Vulkan objects. This extension exposes the actual underlying device memory usage, including allocations that are not normally visible to the application, such as memory consumed by [vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateGraphicsPipelines.html). It is intended primarily for use by debug tooling rather than for production applications.

## [](#_new_structures)New Structures

- [VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportCallbackDataEXT.html)
- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDeviceMemoryReportFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceMemoryReportFeaturesEXT.html)

## [](#_new_function_pointers)New Function Pointers

- [PFN\_vkDeviceMemoryReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDeviceMemoryReportCallbackEXT.html)

## [](#_new_enums)New Enums

- [VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportEventTypeEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDeviceMemoryReportFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEVICE_MEMORY_REPORT_EXTENSION_NAME`
- `VK_EXT_DEVICE_MEMORY_REPORT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_DEVICE_MEMORY_REPORT_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DEVICE_MEMORY_REPORT_CALLBACK_DATA_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_MEMORY_REPORT_FEATURES_EXT`

## [](#_issues)Issues

1\) Should this be better expressed as an extension to VK\_EXT\_debug\_utils and its general-purpose messenger construct?

**RESOLVED**: No. The intended lifecycle is quite different. We want to make this extension tied to the device’s lifecycle. Each ICD just handles its own implementation of this extension, and this extension will only be directly exposed from the ICD. So we can avoid the extra implementation complexity used to accommodate the flexibility of `VK_EXT_debug_utils` extension.

2\) Can we extend and use the existing internal allocation callbacks instead of adding the new callback structure in this extension?

**RESOLVED**: No. Our memory reporting layer that combines this information with other memory information it collects directly (e.g. bindings of resources to [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html)) would have to intercept all entry points that take a [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) parameter and inject its own `pfnInternalAllocation` and `pfnInternalFree`. That may be doable for the extensions we know about, but not for ones we do not. The proposal would work fine in the face of most unknown extensions. But even for ones we know about, since apps can provide a different set of callbacks and userdata and those can be retained by the driver and used later (esp. for pool object, but not just those), we would have to dynamically allocate the interception trampoline every time. That is getting to be an unreasonably large amount of complexity and (possibly) overhead.

We are interested in both alloc/free and import/unimport. The latter is fairly important for tracking (and avoiding double-counting) of swapchain images (still true with “native swapchains” based on external memory) and media/camera interop. Though we might be able to handle this with additional [VkInternalAllocationType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInternalAllocationType.html) values, for import/export we do want to be able to tie this to the external resource, which is one thing that the `memoryObjectId` is for.

The internal alloc/free callbacks are not extensible except via new [VkInternalAllocationType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInternalAllocationType.html) values. The [VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportCallbackDataEXT.html) in this extension is extensible. That was deliberate: there is a real possibility we will want to get extra information in the future. As one example, currently this reports only physical allocations, but we believe there are interesting cases for tracking how populated that VA region is.

The callbacks are clearly specified as only callable within the context of a call from the application into Vulkan. We believe there are some cases where drivers can allocate device memory asynchronously. This was one of the sticky issues that derailed the internal device memory allocation reporting design (which is essentially what this extension is trying to do) leading up to 1.0.

[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) is described in a section called “Host memory” and the intro to it is very explicitly about host memory. The other callbacks are all inherently about host memory. But this extension is very focused on device memory.

3\) Should the callback be reporting which heap is used?

**RESOLVED**: Yes. It is important for non-UMA systems to have all the device memory allocations attributed to the corresponding device memory heaps. For internally-allocated device memory, `heapIndex` will always correspond to an advertised heap, rather than having a magic value indicating a non-advertised heap. Drivers can advertise heaps that do not have any corresponding memory types if they need to.

4\) Should we use an array of callback for the layers to intercept instead of chaining multiple of the [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html) structures in the `pNext` of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html)?

**RESOLVED** No. The pointer to the [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html) structure itself is const and you cannot just cast it away. Thus we cannot update the callback array inside the structure. In addition, we cannot drop this `pNext` chain either, so making a copy of this whole structure does not work either.

5\) Should we track bulk allocations shared among multiple objects?

**RESOLVED** No. Take the shader heap as an example. Some implementations will let multiple [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) objects share the same shader heap. We are not asking the implementation to report `VK_OBJECT_TYPE_PIPELINE` along with a [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) for this bulk allocation. Instead, this bulk allocation is considered as a layer below what this extension is interested in. Later, when the actual [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) objects are created by suballocating from the bulk allocation, we ask the implementation to report the valid handles of the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) objects along with the actual suballocated sizes and different `memoryObjectId`.

6\) Can we require the callbacks to be always called in the same thread with the Vulkan commands?

**RESOLVED** No. Some implementations might choose to multiplex work from multiple application threads into a single backend thread and perform JIT allocations as a part of that flow. Since this behavior is theoretically legit, we cannot require the callbacks to be always called in the same thread with the Vulkan commands, and the note is to remind the applications to handle this case properly.

7\) Should we add an additional “allocation failed” event type with things like size and heap index reported?

**RESOLVED** Yes. This fits in well with the callback infrastructure added in this extension, and implementation touches the same code and has the same overheads as the rest of the extension. It could help debugging things like getting a `VK_ERROR_OUT_OF_HOST_MEMORY` error when ending a command buffer. Right now the allocation failure could have happened anywhere during recording, and a callback would be really useful to understand where and why.

## [](#_version_history)Version History

- Revision 1, 2020-08-26 (Yiwei Zhang)
  
  - Initial version
- Revision 2, 2021-01-06 (Yiwei Zhang)
  
  - Minor description update

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_device_memory_report).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0