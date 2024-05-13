# VK_EXT_texture_compression_astc_hdr(3) Manual Page

## Name

VK_EXT_texture_compression_astc_hdr - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

67

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_texture_compression_astc_hdr%5D%20@janharaldfredriksen-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_texture_compression_astc_hdr%20extension*"
  target="_blank"
  rel="nofollow noopener"><em></em>janharaldfredriksen-arm</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-05-28

**IP Status**  
No known issues.

**Contributors**  
- Jan-Harald Fredriksen, Arm

## <a href="#_description" class="anchor"></a>Description

This extension adds support for textures compressed using the Adaptive
Scalable Texture Compression (ASTC) High Dynamic Range (HDR) profile.

When this extension is enabled, the HDR profile is supported for all
ASTC formats listed in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#appendix-compressedtex-astc"
target="_blank" rel="noopener">ASTC Compressed Image Formats</a>.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_TEXTURE_COMPRESSION_ASTC_HDR_EXTENSION_NAME`

- `VK_EXT_TEXTURE_COMPRESSION_ASTC_HDR_SPEC_VERSION`

- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html):

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

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXTURE_COMPRESSION_ASTC_HDR_FEATURES_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

This extension has been partially promoted. Functionality in this
extension is included in core Vulkan 1.3, with the EXT suffix omitted.
However, the feature is made optional in Vulkan 1.3. The original type,
enum and command names are still available as aliases of the core
functionality.

## <a href="#_issues" class="anchor"></a>Issues

1\) Should we add a feature or limit for this functionality?

Yes. It is consistent with the ASTC LDR support to add a feature like
textureCompressionASTC_HDR.

The feature is strictly speaking redundant as long as this is just an
extension; it would be sufficient to just enable the extension. But
adding the feature is more forward-looking if wanted to make this an
optional core feature in the future.

2\) Should we introduce new format enums for HDR?

Yes. Vulkan 1.0 describes the ASTC format enums as UNORM, e.g.
`VK_FORMAT_ASTC_4x4_UNORM_BLOCK`, so it is confusing to make these
contain HDR data. Note that the OpenGL (ES) extensions did not make this
distinction because a single ASTC HDR texture may contain both unorm and
float blocks. Implementations **may** not be able to distinguish between
LDR and HDR ASTC textures internally and just treat them as the same
format, i.e. if this extension is supported then sampling from a
`VK_FORMAT_ASTC_4x4_UNORM_BLOCK` image format **may** return HDR
results. Applications **can** get predictable results by using the
appropriate image format.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-05-28 (Jan-Harald Fredriksen)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_texture_compression_astc_hdr"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
