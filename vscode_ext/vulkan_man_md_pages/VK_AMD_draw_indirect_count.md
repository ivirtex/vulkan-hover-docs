# VK_AMD_draw_indirect_count(3) Manual Page

## Name

VK_AMD_draw_indirect_count - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

34

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to
  [VK_KHR_draw_indirect_count](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_draw_indirect_count.html)
  extension

  - Which in turn was *promoted* to <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
    target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_draw_indirect_count%5D%20@drakos-amd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_draw_indirect_count%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>drakos-amd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-08-23

**IP Status**  
No known IP claims.

**Contributors**  
- Matthaeus G. Chajdas, AMD

- Derrick Owens, AMD

- Graham Sellers, AMD

- Daniel Rakos, AMD

- Dominik Witczak, AMD

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to source the number of draws for
indirect drawing commands from a buffer. This enables applications to
generate an arbitrary number of drawing commands and execute them
without host intervention.

## <a href="#_promotion_to_vk_khr_draw_indirect_count" class="anchor"></a>Promotion to `VK_KHR_draw_indirect_count`

All functionality in this extension is included in
[`VK_KHR_draw_indirect_count`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_draw_indirect_count.html), with
the suffix changed to KHR. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdDrawIndexedIndirectCountAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirectCountAMD.html)

- [vkCmdDrawIndirectCountAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirectCountAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_DRAW_INDIRECT_COUNT_EXTENSION_NAME`

- `VK_AMD_DRAW_INDIRECT_COUNT_SPEC_VERSION`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2016-08-23 (Dominik Witczak)

  - Minor fixes

- Revision 1, 2016-07-21 (Matthaeus Chajdas)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_draw_indirect_count"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
