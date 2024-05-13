# VK_IMG_format_pvrtc(3) Manual Page

## Name

VK_IMG_format_pvrtc - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

55

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* without replacement

## <a href="#_contact" class="anchor"></a>Contact

- Stuart Smith

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-09-02

**IP Status**  
Imagination Technologies Proprietary

**Contributors**  
- Stuart Smith, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

`VK_IMG_format_pvrtc` provides additional texture compression
functionality specific to Imagination Technologies PowerVR Texture
compression format (called PVRTC).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>As also noted in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#data-format"
target="_blank" rel="noopener">Khronos Data Format Specification</a>,
PVRTC1 images must have dimensions that are a power of two.</p></td>
</tr>
</tbody>
</table>

## <a href="#_deprecation" class="anchor"></a>Deprecation

Both PVRTC1 and PVRTC2 are slower than standard image formats on PowerVR
GPUs, and support will be removed from future hardware.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_IMG_FORMAT_PVRTC_EXTENSION_NAME`

- `VK_IMG_FORMAT_PVRTC_SPEC_VERSION`

- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html):

  - `VK_FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG`

  - `VK_FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG`

  - `VK_FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG`

  - `VK_FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG`

  - `VK_FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG`

  - `VK_FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG`

  - `VK_FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG`

  - `VK_FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-09-02 (Stuart Smith)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_IMG_format_pvrtc"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
