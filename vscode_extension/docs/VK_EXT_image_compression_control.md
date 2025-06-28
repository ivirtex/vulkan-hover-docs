# VK\_EXT\_image\_compression\_control(3) Manual Page

## Name

VK\_EXT\_image\_compression\_control - device extension



## [](#_registered_extension_number)Registered Extension Number

339

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_compression_control%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_compression_control%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_image\_compression\_control](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_image_compression_control.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

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

## [](#_description)Description

This extension enables fixed-rate image compression and adds the ability to control when this kind of compression can be applied. Many implementations support some form of framebuffer compression. This is typically transparent to applications as lossless compression schemes are used. With fixed-rate compression, the compression is done at a defined bitrate. Such compression algorithms generally produce results that are visually lossless, but the results are typically not bit-exact when compared to a non-compressed result. The implementation may not be able to use the requested compression rate in all cases. This extension adds a query that can be used to determine the compression scheme and rate that was applied to an image.

## [](#_new_commands)New Commands

- [vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2EXT.html)

## [](#_new_structures)New Structures

- [VkImageSubresource2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2EXT.html)
- [VkSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2EXT.html)
- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html), [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html):
  
  - [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html)
- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html), [VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFormat2KHR.html), [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html):
  
  - [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionPropertiesEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImageCompressionControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageCompressionControlFeaturesEXT.html)

## [](#_new_enums)New Enums

- [VkImageCompressionFixedRateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFixedRateFlagBitsEXT.html)
- [VkImageCompressionFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFixedRateFlagsEXT.html)
- [VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_IMAGE_COMPRESSION_CONTROL_EXTENSION_NAME`
- `VK_EXT_IMAGE_COMPRESSION_CONTROL_SPEC_VERSION`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_COMPRESSION_EXHAUSTED_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGE_COMPRESSION_CONTROL_EXT`
  - `VK_STRUCTURE_TYPE_IMAGE_COMPRESSION_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_IMAGE_SUBRESOURCE_2_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_COMPRESSION_CONTROL_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_SUBRESOURCE_LAYOUT_2_EXT`

## [](#_version_history)Version History

- Revision 1, 2022-05-02 (Jan-Harald Fredriksen)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_image_compression_control)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0