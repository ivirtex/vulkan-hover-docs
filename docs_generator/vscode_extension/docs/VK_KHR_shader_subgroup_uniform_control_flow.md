# VK\_KHR\_shader\_subgroup\_uniform\_control\_flow(3) Manual Page

## Name

VK\_KHR\_shader\_subgroup\_uniform\_control\_flow - device extension



## [](#_registered_extension_number)Registered Extension Number

324

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_subgroup\_uniform\_control\_flow](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_subgroup_uniform_control_flow.html)

## [](#_contact)Contact

- Alan Baker [\[GitHub\]alan-baker](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_subgroup_uniform_control_flow%5D%20%40alan-baker%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_subgroup_uniform_control_flow%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-08-27

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Requires SPIR-V 1.3.
- This extension provides API support for [`GL_EXT_subgroupuniform_qualifier`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_subgroupuniform_qualifier.txt)

**Contributors**

- Alan Baker, Google
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension allows the use of the `SPV_KHR_subgroup_uniform_control_flow` SPIR-V extension in shader modules. `SPV_KHR_subgroup_uniform_control_flow` provides stronger guarantees that diverged subgroups will reconverge.

Developers should utilize this extension if they use subgroup operations to reduce the work performed by a uniform subgroup. This extension will guarantee that uniform subgroup will reconverge in the same manner as invocation groups (see “Uniform Control Flow” in the [Khronos SPIR-V Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirv-spec)).

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_EXTENSION_NAME`
- `VK_KHR_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_FEATURES_KHR`

## [](#_version_history)Version History

- Revision 1, 2020-08-27 (Alan Baker)
  
  - Internal draft version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_subgroup_uniform_control_flow)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0