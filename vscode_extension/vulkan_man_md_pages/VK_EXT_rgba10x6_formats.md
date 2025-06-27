# VK_EXT_rgba10x6_formats(3) Manual Page

## Name

VK_EXT_rgba10x6_formats - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

345

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_sampler_ycbcr_conversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_rgba10x6_formats%5D%20@janharaldfredriksen-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_rgba10x6_formats%20extension*"
  target="_blank"
  rel="nofollow noopener"><em></em>janharaldfredriksen-arm</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-09-29

**IP Status**  
No known IP claims.

**Contributors**  
- Jan-Harald Fredriksen, Arm

- Graeme Leese, Broadcom

- Spencer Fricke, Samsung

## <a href="#_description" class="anchor"></a>Description

This extension enables the
`VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16` format to be used without
a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
conversion</a> enabled.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_RGBA10X6_FORMATS_EXTENSION_NAME`

- `VK_EXT_RGBA10X6_FORMATS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RGBA10X6_FORMATS_FEATURES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should we reuse the existing format enumeration or introduce a new
one?

**RESOLVED**: We reuse an existing format enumeration,
`VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16`, that was previously
exclusively used for YCbCr and therefore had a set of limitations
related to that usage. The alternative was to introduce a new format
token with exactly the same bit representation as the existing token,
but without the limitations.

2\) Should we only introduce
`VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16` or also 1-3 component
variations?

**RESOLVED**: Only the 4-component format is introduced because the 1-
and 2- component variations are already not exclusive to YCbCr, and the
3-component variation is not a good match for hardware capabilities.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-09-29 (Jan-Harald Fredriksen)

  - Initial EXT version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_rgba10x6_formats"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
