# VK_KHR_shader_subgroup_uniform_control_flow(3) Manual Page

## Name

VK_KHR_shader_subgroup_uniform_control_flow - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

324

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_subgroup_uniform_control_flow](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_subgroup_uniform_control_flow.html)

## <a href="#_contact" class="anchor"></a>Contact

- Alan Baker <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_subgroup_uniform_control_flow%5D%20@alan-baker%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_subgroup_uniform_control_flow%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>alan-baker</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-08-27

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- Requires SPIR-V 1.3.

- This extension provides API support for
  [`GL_EXT_subgroupuniform_qualifier`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_subgroupuniform_qualifier.txt)

**Contributors**  
- Alan Baker, Google

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows the use of the
`SPV_KHR_subgroup_uniform_control_flow` SPIR-V extension in shader
modules. `SPV_KHR_subgroup_uniform_control_flow` provides stronger
guarantees that diverged subgroups will reconverge.

Developers should utilize this extension if they use subgroup operations
to reduce the work performed by a uniform subgroup. This extension will
guarantee that uniform subgroup will reconverge in the same manner as
invocation groups (see “Uniform Control Flow” in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirv-spec"
target="_blank" rel="noopener">Khronos SPIR-V Specification</a>).

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_EXTENSION_NAME`

- `VK_KHR_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_FEATURES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-08-27 (Alan Baker)

  - Internal draft version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_subgroup_uniform_control_flow"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
