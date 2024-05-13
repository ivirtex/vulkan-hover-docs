# VK_EXT_shader_stencil_export(3) Manual Page

## Name

VK_EXT_shader_stencil_export - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

141

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_shader_stencil_export](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_shader_stencil_export.html)

## <a href="#_contact" class="anchor"></a>Contact

- Dominik Witczak <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_stencil_export%5D%20@dominikwitczakamd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_stencil_export%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dominikwitczakamd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-07-19

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARB_shader_stencil_export`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_stencil_export.txt)

**Contributors**  
- Dominik Witczak, AMD

- Daniel Rakos, AMD

- Rex Xu, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the SPIR-V extension
`SPV_EXT_shader_stencil_export`, providing a mechanism whereby a shader
may generate the stencil reference value per invocation. When stencil
testing is enabled, this allows the test to be performed against the
value generated in the shader.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SHADER_STENCIL_EXPORT_EXTENSION_NAME`

- `VK_EXT_SHADER_STENCIL_EXPORT_SPEC_VERSION`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-07-19 (Dominik Witczak)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_shader_stencil_export"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
