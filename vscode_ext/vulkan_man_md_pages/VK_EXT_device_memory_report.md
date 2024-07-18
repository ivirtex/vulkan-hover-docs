# VK_EXT_device_memory_report(3) Manual Page

## Name

VK_EXT_device_memory_report - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

285

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Developer tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Yiwei Zhang <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_device_memory_report%5D%20@zhangyiwei%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_device_memory_report%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>zhangyiwei</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-01-06

**IP Status**  
No known IP claims.

**Contributors**  
- Yiwei Zhang, Google

- Jesse Hall, Google

## <a href="#_description" class="anchor"></a>Description

This device extension allows registration of device memory event
callbacks upon device creation, so that applications or middleware can
obtain detailed information about memory usage and how memory is
associated with Vulkan objects. This extension exposes the actual
underlying device memory usage, including allocations that are not
normally visible to the application, such as memory consumed by
[vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html). It is
intended primarily for use by debug tooling rather than for production
applications.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportCallbackDataEXT.html)

- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDeviceMemoryReportFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceMemoryReportFeaturesEXT.html)

## <a href="#_new_function_pointers" class="anchor"></a>New Function Pointers

- [PFN_vkDeviceMemoryReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDeviceMemoryReportCallbackEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportEventTypeEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkDeviceMemoryReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DEVICE_MEMORY_REPORT_EXTENSION_NAME`

- `VK_EXT_DEVICE_MEMORY_REPORT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_DEVICE_MEMORY_REPORT_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_DEVICE_MEMORY_REPORT_CALLBACK_DATA_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_MEMORY_REPORT_FEATURES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should this be better expressed as an extension to
VK_EXT_debug_utils and its general-purpose messenger construct?

**RESOLVED**: No. The intended lifecycle is quite different. We want to
make this extension tied to the device’s lifecycle. Each ICD just
handles its own implementation of this extension, and this extension
will only be directly exposed from the ICD. So we can avoid the extra
implementation complexity used to accommodate the flexibility of
[`VK_EXT_debug_utils`](VK_EXT_debug_utils.html) extension.

2\) Can we extend and use the existing internal allocation callbacks
instead of adding the new callback structure in this extension?

**RESOLVED**: No. Our memory reporting layer that combines this
information with other memory information it collects directly (e.g.
bindings of resources to [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)) would
have to intercept all entry points that take a
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html) parameter and inject
its own `pfnInternalAllocation` and `pfnInternalFree`. That may be
doable for the extensions we know about, but not for ones we do not. The
proposal would work fine in the face of most unknown extensions. But
even for ones we know about, since apps can provide a different set of
callbacks and userdata and those can be retained by the driver and used
later (esp. for pool object, but not just those), we would have to
dynamically allocate the interception trampoline every time. That is
getting to be an unreasonably large amount of complexity and (possibly)
overhead.

We are interested in both alloc/free and import/unimport. The latter is
fairly important for tracking (and avoiding double-counting) of
swapchain images (still true with “native swapchains” based on external
memory) and media/camera interop. Though we might be able to handle this
with additional
[VkInternalAllocationType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInternalAllocationType.html) values, for
import/export we do want to be able to tie this to the external
resource, which is one thing that the `memoryObjectId` is for.

The internal alloc/free callbacks are not extensible except via new
[VkInternalAllocationType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInternalAllocationType.html) values. The
[VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportCallbackDataEXT.html)
in this extension is extensible. That was deliberate: there is a real
possibility we will want to get extra information in the future. As one
example, currently this reports only physical allocations, but we
believe there are interesting cases for tracking how populated that VA
region is.

The callbacks are clearly specified as only callable within the context
of a call from the application into Vulkan. We believe there are some
cases where drivers can allocate device memory asynchronously. This was
one of the sticky issues that derailed the internal device memory
allocation reporting design (which is essentially what this extension is
trying to do) leading up to 1.0.

[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html) is described in a
section called “Host memory” and the intro to it is very explicitly
about host memory. The other callbacks are all inherently about host
memory. But this extension is very focused on device memory.

3\) Should the callback be reporting which heap is used?

**RESOLVED**: Yes. It is important for non-UMA systems to have all the
device memory allocations attributed to the corresponding device memory
heaps. For internally-allocated device memory, `heapIndex` will always
correspond to an advertised heap, rather than having a magic value
indicating a non-advertised heap. Drivers can advertise heaps that do
not have any corresponding memory types if they need to.

4\) Should we use an array of callback for the layers to intercept
instead of chaining multiple of the
[VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html)
structures in the `pNext` of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html)?

**RESOLVED** No. The pointer to the
[VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html)
structure itself is const and you cannot just cast it away. Thus we
cannot update the callback array inside the structure. In addition, we
cannot drop this `pNext` chain either, so making a copy of this whole
structure does not work either.

5\) Should we track bulk allocations shared among multiple objects?

**RESOLVED** No. Take the shader heap as an example. Some
implementations will let multiple [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) objects
share the same shader heap. We are not asking the implementation to
report `VK_OBJECT_TYPE_PIPELINE` along with a
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) for this bulk allocation. Instead,
this bulk allocation is considered as a layer below what this extension
is interested in. Later, when the actual [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
objects are created by suballocating from the bulk allocation, we ask
the implementation to report the valid handles of the
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) objects along with the actual suballocated
sizes and different `memoryObjectId`.

6\) Can we require the callbacks to be always called in the same thread
with the Vulkan commands?

**RESOLVED** No. Some implementations might choose to multiplex work
from multiple application threads into a single backend thread and
perform JIT allocations as a part of that flow. Since this behavior is
theoretically legit, we cannot require the callbacks to be always called
in the same thread with the Vulkan commands, and the note is to remind
the applications to handle this case properly.

7\) Should we add an additional “allocation failed” event type with
things like size and heap index reported?

**RESOLVED** Yes. This fits in well with the callback infrastructure
added in this extension, and implementation touches the same code and
has the same overheads as the rest of the extension. It could help
debugging things like getting a `VK_ERROR_OUT_OF_HOST_MEMORY` error when
ending a command buffer. Right now the allocation failure could have
happened anywhere during recording, and a callback would be really
useful to understand where and why.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-08-26 (Yiwei Zhang)

  - Initial version

- Revision 2, 2021-01-06 (Yiwei Zhang)

  - Minor description update

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_device_memory_report"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
