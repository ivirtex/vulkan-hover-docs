# VK_EXT_image_compression_control_swapchain(3) Manual Page

## Name

VK_EXT_image_compression_control_swapchain - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

438

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_image_compression_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_compression_control.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_compression_control_swapchain%5D%20@janharaldfredriksen-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_compression_control_swapchain%20extension*"
  target="_blank"
  rel="nofollow noopener"><em></em>janharaldfredriksen-arm</a>

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

- Ian Elliott, Google

## <a href="#_description" class="anchor"></a>Description

This extension enables fixed-rate image compression and adds the ability
to control when this kind of compression can be applied to swapchain
images.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_IMAGE_COMPRESSION_CONTROL_SWAPCHAIN_EXTENSION_NAME`

- `VK_EXT_IMAGE_COMPRESSION_CONTROL_SWAPCHAIN_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_COMPRESSION_CONTROL_SWAPCHAIN_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-05-02 (Jan-Harald Fredriksen)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_compression_control_swapchain"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
