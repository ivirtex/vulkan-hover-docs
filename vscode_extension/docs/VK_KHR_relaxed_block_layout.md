# VK_KHR_relaxed_block_layout(3) Manual Page

## Name

VK_KHR_relaxed_block_layout - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

145

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- John Kessenich <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_relaxed_block_layout%5D%20@johnkslang%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_relaxed_block_layout%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>johnkslang</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-03-26

**IP Status**  
No known IP claims.

**Contributors**  
- John Kessenich, Google

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_relaxed_block_layout` extension allows implementations to
indicate they can support more variation in block `Offset` decorations.
For example, placing a vector of three floats at an offset of 16×N + 4.

See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-resources-layout"
target="_blank" rel="noopener">Offset and Stride Assignment</a> for
details.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_RELAXED_BLOCK_LAYOUT_EXTENSION_NAME`

- `VK_KHR_RELAXED_BLOCK_LAYOUT_SPEC_VERSION`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-03-26 (JohnK)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_relaxed_block_layout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
