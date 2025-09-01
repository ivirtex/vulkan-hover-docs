# VK\_ANDROID\_external\_memory\_android\_hardware\_buffer(3) Manual Page

## Name

VK\_ANDROID\_external\_memory\_android\_hardware\_buffer - device extension



## [](#_registered_extension_number)Registered Extension Number

130

## [](#_revision)Revision

5

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
         and  
         [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
         and  
         [VK\_KHR\_dedicated\_allocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dedicated_allocation.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_format\_feature\_flags2

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ANDROID_external_memory_android_hardware_buffer%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ANDROID_external_memory_android_hardware_buffer%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-09-30

**IP Status**

No known IP claims.

**Contributors**

- Ray Smith, ARM
- Lina Versace, Google
- Jesse Hall, Google
- Tobias Hector, Imagination
- James Jones, NVIDIA
- Tony Zlatinski, NVIDIA
- Matthew Netsch, Qualcomm
- Andrew Garrard, Samsung

## [](#_description)Description

This extension enables an application to import Android [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) objects created outside of the Vulkan device into Vulkan memory objects, where they **can** be bound to images and buffers. It also allows exporting an [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) from a Vulkan memory object for symmetry with other operating systems. But since not all [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) usages and formats have Vulkan equivalents, exporting from Vulkan provides strictly less functionality than creating the [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) externally and importing it.

Some [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) images have implementation-defined *external formats* that **may** not correspond to Vulkan formats. Sampler Y′CBCR conversion **can** be used to sample from these images and convert them to a known color space.

## [](#_new_base_types)New Base Types

- [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html)

## [](#_new_commands)New Commands

- [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)
- [vkGetMemoryAndroidHardwareBufferANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryAndroidHardwareBufferANDROID.html)

## [](#_new_structures)New Structures

- [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferPropertiesANDROID.html)
- [VkMemoryGetAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html)
- Extending [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferPropertiesANDROID.html):
  
  - [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html)
- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html), [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription2.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html):
  
  - [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)
- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html):
  
  - [VkAndroidHardwareBufferUsageANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferUsageANDROID.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)

If [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferPropertiesANDROID.html):
  
  - [VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ANDROID_EXTERNAL_MEMORY_ANDROID_HARDWARE_BUFFER_EXTENSION_NAME`
- `VK_ANDROID_EXTERNAL_MEMORY_ANDROID_HARDWARE_BUFFER_SPEC_VERSION`
- Extending [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID`
  - `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID`
  - `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID`
  - `VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID`
  - `VK_STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID`
  - `VK_STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID`

If [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_2_ANDROID`

## [](#_issues)Issues

1\) Other external memory objects are represented as weakly-typed handles (e.g. Win32 `HANDLE` or POSIX file descriptor), and require a handle type parameter along with handles. [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) is strongly typed, so naming the handle type is redundant. Does symmetry justify adding handle type parameters/fields anyway?

**RESOLVED**: No. The handle type is already provided in places that treat external memory objects generically. In the places we would add it, the application code that would have to provide the handle type value is already dealing with [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html)-specific commands/structures; the extra symmetry would not be enough to make that code generic.

2\) The internal layout and therefore size of a [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) image may depend on native usage flags that do not have corresponding Vulkan counterparts. Do we provide this information to [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html) somehow, or allow the allocation size reported by [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements.html) to be approximate?

**RESOLVED**: Allow the allocation size to be unspecified when allocating the memory. It has to work this way for exported image memory anyway, since [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) allocation happens in [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateMemory.html), and internally is performed by a separate HAL, not the Vulkan implementation itself. There is a similar issue with [vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout.html): the layout is determined by the allocator HAL, so it is not known until the image is bound to memory.

3\) Should the result of sampling an external-format image with the suggested Y′CBCR conversion parameters yield the same results as using a `samplerExternalOES` in OpenGL ES?

**RESOLVED**: This would be desirable, so that apps converting from OpenGL ES to Vulkan could get the same output given the same input. But since sampling and conversion from Y′CBCR images is so loosely defined in OpenGL ES, multiple implementations do it in a way that does not conform to Vulkan’s requirements. Modifying the OpenGL ES implementation would be difficult, and would change the output of existing unmodified applications. Changing the output only for applications that are being modified gives developers the chance to notice and mitigate any problems. Implementations are encouraged to minimize differences as much as possible without causing compatibility problems for existing OpenGL ES applications or violating Vulkan requirements.

4\) Should an [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) with `AHARDWAREBUFFER_USAGE_CPU_*` usage be mappable in Vulkan? Should it be possible to export an `AHardwareBuffers` with such usage?

**RESOLVED**: Optional, and mapping in Vulkan is not the same as `AHardwareBuffer_lock`. The semantics of these are different: mapping in memory is persistent, just gives a raw view of the memory contents, and does not involve ownership. `AHardwareBuffer_lock` gives the host exclusive access to the buffer, is temporary, and allows for reformatting copy-in/copy-out. Implementations are not required to support host-visible memory types for imported Android hardware buffers or resources backed by them. If a host-visible memory type is supported and used, the memory can be mapped in Vulkan, but doing so follows Vulkan semantics: it is just a raw view of the data and does not imply ownership (this means implementations must not internally call `AHardwareBuffer_lock` to implement [vkMapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory.html), or assume the application has done so). Implementations are not required to support linear-tiled images backed by Android hardware buffers, even if the [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) has CPU usage. There is no reliable way to allocate memory in Vulkan that can be exported to a [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) with CPU usage.

5\) Android may add new [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) formats and usage flags over time. Can reference to them be added to this extension, or do they need a new extension?

**RESOLVED**: This extension can document the interaction between the new AHB formats/usages and existing Vulkan features. No new Vulkan features or implementation requirements can be added. The extension version number will be incremented when this additional documentation is added, but the version number does not indicate that an implementation supports Vulkan memory or resources that map to the new [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) features: support for that must be queried with [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) or is implied by successfully allocating a [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) outside of Vulkan that uses the new feature and has a GPU usage flag.

In essence, these are new features added to a new Android API level, rather than new Vulkan features. The extension will only document how existing Vulkan features map to that new Android feature.

## [](#_version_history)Version History

- Revision 5, 2022-02-04 (Chris Forbes)
  
  - Describe mapping of flags for storage image support
- Revision 4, 2021-09-30 (Jon Leech)
  
  - Add interaction with `VK_KHR_format_feature_flags2` to `vk.xml`
- Revision 3, 2019-08-27 (Jon Leech)
  
  - Update revision history to correspond to XML version number
- Revision 2, 2018-04-09 (Petr Kraus)
  
  - Markup fixes and remove incorrect Draft status
- Revision 1, 2018-03-04 (Jesse Hall)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ANDROID_external_memory_android_hardware_buffer).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0