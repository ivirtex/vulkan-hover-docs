# VK\_EXT\_texture\_compression\_astc\_hdr(3) Manual Page

## Name

VK\_EXT\_texture\_compression\_astc\_hdr - device extension



## [](#_registered_extension_number)Registered Extension Number

67

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_texture_compression_astc_hdr%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_texture_compression_astc_hdr%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-05-28

**IP Status**

No known issues.

**Contributors**

- Jan-Harald Fredriksen, Arm

## [](#_description)Description

This extension adds support for textures compressed using the Adaptive Scalable Texture Compression (ASTC) High Dynamic Range (HDR) profile.

When this extension is enabled, the HDR profile is supported for all ASTC formats listed in [ASTC Compressed Image Formats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#appendix-compressedtex-astc).

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_TEXTURE_COMPRESSION_ASTC_HDR_EXTENSION_NAME`
- `VK_EXT_TEXTURE_COMPRESSION_ASTC_HDR_SPEC_VERSION`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_ASTC_10x10_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_10x5_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_10x6_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_10x8_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_12x10_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_12x12_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_4x4_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_5x4_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_5x5_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_6x5_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_6x6_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_8x5_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_8x6_SFLOAT_BLOCK_EXT`
  - `VK_FORMAT_ASTC_8x8_SFLOAT_BLOCK_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXTURE_COMPRESSION_ASTC_HDR_FEATURES_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. However, the feature is made optional in Vulkan 1.3. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_issues)Issues

1\) Should we add a feature or limit for this functionality?

Yes. It is consistent with the ASTC LDR support to add a feature like textureCompressionASTC\_HDR.

The feature is strictly speaking redundant as long as this is just an extension; it would be sufficient to just enable the extension. But adding the feature is more forward-looking if wanted to make this an optional core feature in the future.

2\) Should we introduce new format enums for HDR?

Yes. Vulkan 1.0 describes the ASTC format enums as UNORM, e.g. `VK_FORMAT_ASTC_4x4_UNORM_BLOCK`, so it is confusing to make these contain HDR data. Note that the OpenGL (ES) extensions did not make this distinction because a single ASTC HDR texture may contain both unorm and float blocks. Implementations **may** not be able to distinguish between LDR and HDR ASTC textures internally and just treat them as the same format, i.e. if this extension is supported then sampling from a `VK_FORMAT_ASTC_4x4_UNORM_BLOCK` image format **may** return HDR results. Applications **can** get predictable results by using the appropriate image format.

## [](#_version_history)Version History

- Revision 1, 2019-05-28 (Jan-Harald Fredriksen)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_texture_compression_astc_hdr).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0