# VK\_EXT\_shader\_stencil\_export(3) Manual Page

## Name

VK\_EXT\_shader\_stencil\_export - device extension



## [](#_registered_extension_number)Registered Extension Number

141

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_shader\_stencil\_export](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_shader_stencil_export.html)

## [](#_contact)Contact

- Dominik Witczak [\[GitHub\]dominikwitczakamd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_stencil_export%5D%20%40dominikwitczakamd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_stencil_export%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-07-19

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_ARB_shader_stencil_export`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_stencil_export.txt)

**Contributors**

- Dominik Witczak, AMD
- Daniel Rakos, AMD
- Rex Xu, AMD

## [](#_description)Description

This extension adds support for the SPIR-V extension `SPV_EXT_shader_stencil_export`, providing a mechanism whereby a shader may generate the stencil reference value per invocation. When stencil testing is enabled, this allows the test to be performed against the value generated in the shader.

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_STENCIL_EXPORT_EXTENSION_NAME`
- `VK_EXT_SHADER_STENCIL_EXPORT_SPEC_VERSION`

## [](#_version_history)Version History

- Revision 1, 2017-07-19 (Dominik Witczak)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_stencil_export).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0