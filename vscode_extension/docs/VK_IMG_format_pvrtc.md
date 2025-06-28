# VK\_IMG\_format\_pvrtc(3) Manual Page

## Name

VK\_IMG\_format\_pvrtc - device extension



## [](#_registered_extension_number)Registered Extension Number

55

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Deprecated* without replacement

## [](#_contact)Contact

- Stuart Smith

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-09-02

**IP Status**

Imagination Technologies Proprietary

**Contributors**

- Stuart Smith, Imagination Technologies

## [](#_description)Description

`VK_IMG_format_pvrtc` provides additional texture compression functionality specific to Imagination Technologies PowerVR Texture compression format (called PVRTC).

Note

As also noted in the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format), PVRTC1 images must have dimensions that are a power of two.

## [](#_deprecation)Deprecation

Both PVRTC1 and PVRTC2 are slower than standard image formats on PowerVR GPUs, and support will be removed from future hardware.

## [](#_new_enum_constants)New Enum Constants

- `VK_IMG_FORMAT_PVRTC_EXTENSION_NAME`
- `VK_IMG_FORMAT_PVRTC_SPEC_VERSION`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG`
  - `VK_FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG`
  - `VK_FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG`
  - `VK_FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG`
  - `VK_FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG`
  - `VK_FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG`
  - `VK_FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG`
  - `VK_FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG`

## [](#_version_history)Version History

- Revision 1, 2019-09-02 (Stuart Smith)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_IMG_format_pvrtc)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0