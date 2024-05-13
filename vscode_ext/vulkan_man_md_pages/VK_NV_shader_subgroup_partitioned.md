# VK_NV_shader_subgroup_partitioned(3) Manual Page

## Name

VK_NV_shader_subgroup_partitioned - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

199

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_shader_subgroup_partitioned](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_shader_subgroup_partitioned.html)

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_shader_subgroup_partitioned%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_shader_subgroup_partitioned%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-03-17

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_shader_subgroup_partitioned`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GL_NV_shader_subgroup_partitioned.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension enables support for a new class of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-group-operations"
target="_blank" rel="noopener">group operations</a> on <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-subgroup"
target="_blank" rel="noopener">subgroups</a> via the
[`GL_NV_shader_subgroup_partitioned`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GL_NV_shader_subgroup_partitioned.txt)
GLSL extension and
[`SPV_NV_shader_subgroup_partitioned`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_shader_subgroup_partitioned.html)
SPIR-V extension. Support for these new operations is advertised via the
`VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV` bit.

This extension requires Vulkan 1.1, for general subgroup support.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_SHADER_SUBGROUP_PARTITIONED_EXTENSION_NAME`

- `VK_NV_SHADER_SUBGROUP_PARTITIONED_SPEC_VERSION`

- Extending [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html):

  - `VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-03-17 (Jeff Bolz)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_shader_subgroup_partitioned"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
