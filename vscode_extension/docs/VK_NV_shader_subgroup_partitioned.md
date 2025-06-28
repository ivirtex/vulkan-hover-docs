# VK\_NV\_shader\_subgroup\_partitioned(3) Manual Page

## Name

VK\_NV\_shader\_subgroup\_partitioned - device extension



## [](#_registered_extension_number)Registered Extension Number

199

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_shader\_subgroup\_partitioned](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_shader_subgroup_partitioned.html)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_shader_subgroup_partitioned%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_shader_subgroup_partitioned%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-03-17

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_shader_subgroup_partitioned`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GL_NV_shader_subgroup_partitioned.txt)

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension enables support for a new class of [group operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-group-operations) on [subgroups](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-scope-subgroup) via the [`GL_NV_shader_subgroup_partitioned`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GL_NV_shader_subgroup_partitioned.txt) GLSL extension and [`SPV_NV_shader_subgroup_partitioned`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_shader_subgroup_partitioned.html) SPIR-V extension. Support for these new operations is advertised via the `VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV` bit.

This extension requires Vulkan 1.1, for general subgroup support.

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_SHADER_SUBGROUP_PARTITIONED_EXTENSION_NAME`
- `VK_NV_SHADER_SUBGROUP_PARTITIONED_SPEC_VERSION`
- Extending [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubgroupFeatureFlagBits.html):
  
  - `VK_SUBGROUP_FEATURE_PARTITIONED_BIT_NV`

## [](#_version_history)Version History

- Revision 1, 2018-03-17 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_shader_subgroup_partitioned)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0