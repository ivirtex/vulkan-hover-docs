# VK\_ARM\_format\_pack(3) Manual Page

## Name

VK\_ARM\_format\_pack - device extension



## [](#_registered_extension_number)Registered Extension Number

610

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_format_pack%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_format_pack%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-03-24

**Interactions and External Dependencies**

**Contributors**

- Jan-Harald Fredriksen, Arm
- Lisa Wu, Arm
- Oivind Boge, Arm

## [](#_description)Description

This extension adds support for additional 1-, 2- and 4-component formats with 10, 12, or 14 bits of components in 16-bit containers.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFormatPackFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFormatPackFeaturesARM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ARM_FORMAT_PACK_EXTENSION_NAME`
- `VK_ARM_FORMAT_PACK_SPEC_VERSION`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_G14X2_B14X2R14X2_2PLANE_420_UNORM_3PACK16_ARM`
  - `VK_FORMAT_G14X2_B14X2R14X2_2PLANE_422_UNORM_3PACK16_ARM`
  - `VK_FORMAT_R10X6G10X6B10X6A10X6_UINT_4PACK16_ARM`
  - `VK_FORMAT_R10X6G10X6_UINT_2PACK16_ARM`
  - `VK_FORMAT_R10X6_UINT_PACK16_ARM`
  - `VK_FORMAT_R12X4G12X4B12X4A12X4_UINT_4PACK16_ARM`
  - `VK_FORMAT_R12X4G12X4_UINT_2PACK16_ARM`
  - `VK_FORMAT_R12X4_UINT_PACK16_ARM`
  - `VK_FORMAT_R14X2G14X2B14X2A14X2_UINT_4PACK16_ARM`
  - `VK_FORMAT_R14X2G14X2B14X2A14X2_UNORM_4PACK16_ARM`
  - `VK_FORMAT_R14X2G14X2_UINT_2PACK16_ARM`
  - `VK_FORMAT_R14X2G14X2_UNORM_2PACK16_ARM`
  - `VK_FORMAT_R14X2_UINT_PACK16_ARM`
  - `VK_FORMAT_R14X2_UNORM_PACK16_ARM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FORMAT_PACK_FEATURES_ARM`

## [](#_issues)Issues

### [](#_what_do_we_call_this_extension)What do we call this extension?

**RESOLVED**

Many existing extension have the format in the name, but in this case we want to expose multiple formats.

We will describe this set of formats as a "pack".

### [](#_compatibility_classes)Compatibility classes

**RESOLVED**

Should these additional formats be in the same compatibility class as any other formats? For single-plane formats, we put formats with the same number of bits (but different types) in the same class. Each multi-plane or subsampled format gets its own compatibility class. This is consistent with how existing formats are handled.

### [](#_format_feature_requirements)Format feature requirements

**RESOLVED**

The format feature queries could be used to determine what is supported on any given implementation, but it may be useful to establish a baseline requirement in the specification. For that purpose, we require a set of format features - sufficient to enable texture operations - to be supported for the added unsigned integer single-plane formats. Other formats and format features are optional.

## [](#_version_history)Version History

- Revision 1, 2025-03-24
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ARM_format_pack)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0