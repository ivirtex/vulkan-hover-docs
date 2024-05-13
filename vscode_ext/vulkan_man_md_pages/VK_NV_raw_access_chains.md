# VK_NV_raw_access_chains(3) Manual Page

## Name

VK_NV_raw_access_chains - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

556

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_raw_access_chains](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_raw_access_chains.html)

## <a href="#_contact" class="anchor"></a>Contact

- Rodrigo Locatti <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_raw_access_chains%5D%20@rlocatti%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_raw_access_chains%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>rlocatti</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-12-04

**Interactions and External Dependencies**  
- This extension requires
  [`SPV_NV_raw_access_chains`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_raw_access_chains.html)

**Contributors**  
- Hans-Kristian Arntzen, Valve

- Rodrigo Locatti, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows the use of the `SPV_NV_raw_access_chains`
extension in SPIR-V shader modules. This enables SPIR-V producers to
efficiently implement interfaces similar to Direct3D structured buffers
and byte address buffers, allowing shaders compiled from an HLSL source
to generate more efficient code.

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RawAccessChainsNV"
  target="_blank" rel="noopener">RawAccessChainsNV</a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-12-04 (Rodrigo Locatti)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_raw_access_chains"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
