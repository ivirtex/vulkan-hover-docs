# VK\_EXT\_border\_color\_swizzle(3) Manual Page

## Name

VK\_EXT\_border\_color\_swizzle - device extension



## [](#_registered_extension_number)Registered Extension Number

412

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_custom\_border\_color](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_custom_border_color.html)

## [](#_special_uses)Special Uses

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)
- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_border_color_swizzle%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_border_color_swizzle%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-10-12

**IP Status**

No known IP claims.

**Contributors**

- Graeme Leese, Broadcom
- Jan-Harald Fredriksen, Arm
- Ricardo Garcia, Igalia
- Shahbaz Youssefi, Google
- Stu Smith, AMD

## [](#_description)Description

After the publication of VK\_EXT\_custom\_border\_color, it was discovered that some implementations had undefined behavior when combining a sampler that uses a custom border color with image views whose component mapping is not the identity mapping.

Since VK\_EXT\_custom\_border\_color has already shipped, this new extension VK\_EXT\_border\_color\_swizzle was created to define the interaction between custom border colors and non-identity image view swizzles, and provide a work-around for implementations that must pre-swizzle the sampler border color to match the image view component mapping it is combined with.

This extension also defines the behavior between samplers with an opaque black border color and image views with a non-identity component swizzle, which was previously left undefined.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceBorderColorSwizzleFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBorderColorSwizzleFeaturesEXT.html)
- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html):
  
  - [VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_BORDER_COLOR_SWIZZLE_EXTENSION_NAME`
- `VK_EXT_BORDER_COLOR_SWIZZLE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BORDER_COLOR_SWIZZLE_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_SAMPLER_BORDER_COLOR_COMPONENT_MAPPING_CREATE_INFO_EXT`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2021-10-12 (Piers Daniell)
  
  - Internal revisions.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_border_color_swizzle).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0