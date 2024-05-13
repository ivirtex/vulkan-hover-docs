# VK_EXT_host_image_copy(3) Manual Page

## Name

VK_EXT_host_image_copy - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

271

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

        
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html)  
     and  
     [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_host_image_copy%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_host_image_copy%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_host_image_copy](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_host_image_copy.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to copy data between host memory and
images on the host processor, without staging the data through a
GPU-accessible buffer. This removes the need to allocate and manage the
buffer and its associated memory. On some architectures it may also
eliminate an extra copy operation. This extension additionally allows
applications to copy data between images on the host.

To support initializing a new image in preparation for a host copy, it
is now possible to transition a new image to `VK_IMAGE_LAYOUT_GENERAL`
or other host-copyable layouts via
[vkTransitionImageLayoutEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTransitionImageLayoutEXT.html).
Additionally, it is possible to perform copies that preserve the
swizzling layout of the image by using the
`VK_HOST_IMAGE_COPY_MEMCPY_EXT` flag. In that case, the memory size
needed for copies to or from a buffer can be retrieved by chaining
[VkSubresourceHostMemcpySizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceHostMemcpySizeEXT.html) to
`pLayout` in
[vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2EXT.html).

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCopyImageToImageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyImageToImageEXT.html)

- [vkCopyImageToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyImageToMemoryEXT.html)

- [vkCopyMemoryToImageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToImageEXT.html)

- [vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2EXT.html)

- [vkTransitionImageLayoutEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTransitionImageLayoutEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToImageInfoEXT.html)

- [VkCopyImageToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToMemoryInfoEXT.html)

- [VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html)

- [VkHostImageLayoutTransitionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageLayoutTransitionInfoEXT.html)

- [VkImageSubresource2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2EXT.html)

- [VkImageToMemoryCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageToMemoryCopyEXT.html)

- [VkMemoryToImageCopyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryToImageCopyEXT.html)

- [VkSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2EXT.html)

- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html):

  - [VkHostImageCopyDevicePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyDevicePerformanceQueryEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceHostImageCopyFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)

- Extending [VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html):

  - [VkSubresourceHostMemcpySizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceHostMemcpySizeEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkHostImageCopyFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkHostImageCopyFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_HOST_IMAGE_COPY_EXTENSION_NAME`

- `VK_EXT_HOST_IMAGE_COPY_SPEC_VERSION`

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html):

  - `VK_FORMAT_FEATURE_2_HOST_IMAGE_TRANSFER_BIT_EXT`

- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html):

  - `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

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

## <a href="#_issues" class="anchor"></a>Issues

1\) When uploading data to an image, the data is usually loaded from
disk. Why not have the application load the data directly into a
`VkDeviceMemory` bound to a buffer (instead of host memory), and use
[vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html)? The same could be
done when downloading data from an image.

**RESOLVED**: This may not always be possible. Complicated Vulkan
applications such as game engines often have decoupled subsystems for
streaming data and rendering. It may be unreasonable to require the
streaming subsystem to coordinate with the rendering subsystem to
allocate memory on its behalf, especially as Vulkan may not be the only
API supported by the engine. In emulation layers, the image data is
necessarily provided by the application in host memory, so an
optimization as suggested is not possible. Most importantly, the device
memory may not be mappable by an application, but still accessible to
the driver.

2\) Are `optimalBufferCopyOffsetAlignment` and
`optimalBufferCopyRowPitchAlignment` applicable to host memory as well
with the functions introduced by this extension? Or should there be new
limits?

**RESOLVED**: No alignment requirements for the host memory pointer.

3\) Should there be granularity requirements for image offsets and
extents?

**RESOLVED**: No granularity requirements, i.e. a granularity of 1 pixel
(for non-compressed formats) and 1 texel block (for compressed formats)
is assumed.

4\) How should the application deal with layout transitions before or
after copying to or from images?

**RESOLVED**: An existing issue with linear images is that when
emulating other APIs, it is impossible to know when to transition them
as they are written to by the host and then used bindlessly. The copy
operations in this extension are affected by the same limitation. A new
command is thus introduced by this extension to address this problem by
allowing the host to perform an image layout transition between a
handful of layouts.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2021-01-20 (Faith Ekstrand)

  - Initial idea and xml

- Revision 1, 2023-04-26 (Shahbaz Youssefi)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_host_image_copy"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
