# VK\_EXT\_host\_image\_copy(3) Manual Page

## Name

VK\_EXT\_host\_image\_copy - device extension



## [](#_registered_extension_number)Registered Extension Number

271

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html)  
     and  
     [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_host_image_copy%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_host_image_copy%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_host\_image\_copy](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_host_image_copy.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-04-26

**Contributors**

- Shahbaz Youssefi, Google
- Faith Ekstrand, Collabora
- Hans-Kristian Arntzen, Valve
- Piers Daniell, NVIDIA
- Jan-Harald Fredriksen, Arm
- James Fitzpatrick, Imagination
- Daniel Story, Nintendo

## [](#_description)Description

This extension allows applications to copy data between host memory and images on the host processor, without staging the data through a GPU-accessible buffer. This removes the need to allocate and manage the buffer and its associated memory. On some architectures it may also eliminate an extra copy operation. This extension additionally allows applications to copy data between images on the host.

To support initializing a new image in preparation for a host copy, it is now possible to transition a new image to `VK_IMAGE_LAYOUT_GENERAL` or other host-copyable layouts via [vkTransitionImageLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTransitionImageLayoutEXT.html). Additionally, it is possible to perform copies that preserve the swizzling layout of the image by using the `VK_HOST_IMAGE_COPY_MEMCPY_BIT_EXT` flag. In that case, the memory size needed for copies to or from a buffer can be retrieved by chaining [VkSubresourceHostMemcpySizeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceHostMemcpySizeEXT.html) to `pLayout` in [vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2EXT.html).

## [](#_new_commands)New Commands

- [vkCopyImageToImageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyImageToImageEXT.html)
- [vkCopyImageToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyImageToMemoryEXT.html)
- [vkCopyMemoryToImageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToImageEXT.html)
- [vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2EXT.html)
- [vkTransitionImageLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTransitionImageLayoutEXT.html)

## [](#_new_structures)New Structures

- [VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfoEXT.html)
- [VkCopyImageToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfoEXT.html)
- [VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfoEXT.html)
- [VkHostImageLayoutTransitionInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageLayoutTransitionInfoEXT.html)
- [VkImageSubresource2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2EXT.html)
- [VkImageToMemoryCopyEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageToMemoryCopyEXT.html)
- [VkMemoryToImageCopyEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryToImageCopyEXT.html)
- [VkSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2EXT.html)
- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html):
  
  - [VkHostImageCopyDevicePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyDevicePerformanceQueryEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceHostImageCopyFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostImageCopyFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)
- Extending [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html):
  
  - [VkSubresourceHostMemcpySizeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceHostMemcpySizeEXT.html)

## [](#_new_enums)New Enums

- [VkHostImageCopyFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkHostImageCopyFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_HOST_IMAGE_COPY_EXTENSION_NAME`
- `VK_EXT_HOST_IMAGE_COPY_SPEC_VERSION`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_HOST_IMAGE_TRANSFER_BIT_EXT`
- Extending [VkHostImageCopyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlagBits.html):
  
  - `VK_HOST_IMAGE_COPY_MEMCPY_BIT_EXT`
  - `VK_HOST_IMAGE_COPY_MEMCPY_EXT`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_IMAGE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_MEMORY_INFO_EXT`
  - `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_IMAGE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_HOST_IMAGE_COPY_DEVICE_PERFORMANCE_QUERY_EXT`
  - `VK_STRUCTURE_TYPE_HOST_IMAGE_LAYOUT_TRANSITION_INFO_EXT`
  - `VK_STRUCTURE_TYPE_IMAGE_TO_MEMORY_COPY_EXT`
  - `VK_STRUCTURE_TYPE_MEMORY_TO_IMAGE_COPY_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_IMAGE_COPY_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_IMAGE_COPY_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_SUBRESOURCE_HOST_MEMCPY_SIZE_EXT`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4, with the EXT suffix omitted. However, the feature is made optional in Vulkan 1.4. The original type, enum and command names are still available as aliases of the core functionality.

A Vulkan 1.4 implementation that has a `VK_QUEUE_GRAPHICS_BIT` queue must support either:

- the [`hostImageCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-hostImageCopy) feature; or
- an additional queue that supports `VK_QUEUE_TRANSFER_BIT`.

Additionally, all queues supporting `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT` must also advertise `VK_QUEUE_TRANSFER_BIT`.

## [](#_issues)Issues

1\) When uploading data to an image, the data is usually loaded from disk. Why not have the application load the data directly into a `VkDeviceMemory` bound to a buffer (instead of host memory), and use [vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage.html)? The same could be done when downloading data from an image.

**RESOLVED**: This may not always be possible. Complicated Vulkan applications such as game engines often have decoupled subsystems for streaming data and rendering. It may be unreasonable to require the streaming subsystem to coordinate with the rendering subsystem to allocate memory on its behalf, especially as Vulkan may not be the only API supported by the engine. In emulation layers, the image data is necessarily provided by the application in host memory, so an optimization as suggested is not possible. Most importantly, the device memory may not be mappable by an application, but still accessible to the driver.

2\) Are `optimalBufferCopyOffsetAlignment` and `optimalBufferCopyRowPitchAlignment` applicable to host memory as well with the functions introduced by this extension? Or should there be new limits?

**RESOLVED**: No alignment requirements for the host memory pointer.

3\) Should there be granularity requirements for image offsets and extents?

**RESOLVED**: No granularity requirements, i.e. a granularity of 1 pixel (for non-compressed formats) and 1 texel block (for compressed formats) is assumed.

4\) How should the application deal with layout transitions before or after copying to or from images?

**RESOLVED**: An existing issue with linear images is that when emulating other APIs, it is impossible to know when to transition them as they are written to by the host and then used bindlessly. The copy operations in this extension are affected by the same limitation. A new command is thus introduced by this extension to address this problem by allowing the host to perform an image layout transition between a handful of layouts.

## [](#_version_history)Version History

- Revision 0, 2021-01-20 (Faith Ekstrand)
  
  - Initial idea and xml
- Revision 1, 2023-04-26 (Shahbaz Youssefi)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_host_image_copy)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0