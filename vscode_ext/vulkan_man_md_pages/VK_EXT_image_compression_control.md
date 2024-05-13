# VK_EXT_image_compression_control(3) Manual Page

## Name

VK_EXT_image_compression_control - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

339

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_compression_control%5D%20@janharaldfredriksen-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_compression_control%20extension*"
  target="_blank"
  rel="nofollow noopener"><em></em>janharaldfredriksen-arm</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_image_compression_control](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_image_compression_control.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-05-02

**IP Status**  
No known IP claims.

**Contributors**  
- Jan-Harald Fredriksen, Arm

- Graeme Leese, Broadcom

- Andrew Garrard, Imagination

- Lisa Wu, Arm

- Peter Kohaut, Arm

## <a href="#_description" class="anchor"></a>Description

This extension enables fixed-rate image compression and adds the ability
to control when this kind of compression can be applied. Many
implementations support some form of framebuffer compression. This is
typically transparent to applications as lossless compression schemes
are used. With fixed-rate compression, the compression is done at a
defined bitrate. Such compression algorithms generally produce results
that are visually lossless, but the results are typically not bit-exact
when compared to a non-compressed result. The implementation may not be
able to use the requested compression rate in all cases. This extension
adds a query that can be used to determine the compression scheme and
rate that was applied to an image.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2EXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkImageSubresource2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2EXT.html)

- [VkSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2EXT.html)

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html),
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html):

  - [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)

- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html),
  [VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormat2KHR.html),
  [VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html):

  - [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceImageCompressionControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageCompressionControlFeaturesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkImageCompressionFixedRateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFixedRateFlagBitsEXT.html)

- [VkImageCompressionFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFixedRateFlagsEXT.html)

- [VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_IMAGE_COMPRESSION_CONTROL_EXTENSION_NAME`

- `VK_EXT_IMAGE_COMPRESSION_CONTROL_SPEC_VERSION`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_COMPRESSION_EXHAUSTED_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_COMPRESSION_CONTROL_EXT`

  - `VK_STRUCTURE_TYPE_IMAGE_COMPRESSION_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_IMAGE_SUBRESOURCE_2_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_COMPRESSION_CONTROL_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_SUBRESOURCE_LAYOUT_2_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-05-02 (Jan-Harald Fredriksen)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_compression_control"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
