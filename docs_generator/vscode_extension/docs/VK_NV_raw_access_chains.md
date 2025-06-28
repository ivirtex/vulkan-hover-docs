# VK\_NV\_raw\_access\_chains(3) Manual Page

## Name

VK\_NV\_raw\_access\_chains - device extension



## [](#_registered_extension_number)Registered Extension Number

556

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_raw\_access\_chains](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_raw_access_chains.html)

## [](#_contact)Contact

- Rodrigo Locatti [\[GitHub\]rlocatti](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_raw_access_chains%5D%20%40rlocatti%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_raw_access_chains%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-12-04

**Interactions and External Dependencies**

- This extension requires [`SPV_NV_raw_access_chains`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_raw_access_chains.html)

**Contributors**

- Hans-Kristian Arntzen, Valve
- Rodrigo Locatti, NVIDIA

## [](#_description)Description

This extension allows the use of the `SPV_NV_raw_access_chains` extension in SPIR-V shader modules. This enables SPIR-V producers to efficiently implement interfaces similar to Direct3D structured buffers and byte address buffers, allowing shaders compiled from an HLSL source to generate more efficient code.

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [RawAccessChainsNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RawAccessChainsNV)

## [](#_version_history)Version History

- Revision 1, 2023-12-04 (Rodrigo Locatti)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_raw_access_chains)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0