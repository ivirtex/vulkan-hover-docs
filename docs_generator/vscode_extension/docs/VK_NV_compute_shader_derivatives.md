# VK\_NV\_compute\_shader\_derivatives(3) Manual Page

## Name

VK\_NV\_compute\_shader\_derivatives - device extension



## [](#_registered_extension_number)Registered Extension Number

202

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_compute\_shader\_derivatives](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_compute_shader_derivatives.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_compute\_shader\_derivatives](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_compute_shader_derivatives.html) extension

## [](#_contact)Contact

- Pat Brown [\[GitHub\]nvpbrown](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_compute_shader_derivatives%5D%20%40nvpbrown%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_compute_shader_derivatives%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-07-19

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_compute_shader_derivatives`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_compute_shader_derivatives.txt)

**Contributors**

- Pat Brown, NVIDIA

## [](#_description)Description

This extension adds Vulkan support for the [`SPV_NV_compute_shader_derivatives`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_compute_shader_derivatives.html) SPIR-V extension.

The SPIR-V extension provides two new execution modes, both of which allow compute shaders to use built-ins that evaluate compute derivatives explicitly or implicitly. Derivatives will be computed via differencing over a 2x2 group of shader invocations. The `DerivativeGroupQuadsNV` execution mode assembles shader invocations into 2x2 groups, where each group has x and y coordinates of the local invocation ID of the form (2m+{0,1}, 2n+{0,1}). The `DerivativeGroupLinearNV` execution mode assembles shader invocations into 2x2 groups, where each group has local invocation index values of the form 4m+{0,1,2,3}.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceComputeShaderDerivativesFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceComputeShaderDerivativesFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_COMPUTE_SHADER_DERIVATIVES_EXTENSION_NAME`
- `VK_NV_COMPUTE_SHADER_DERIVATIVES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV`

## [](#_new_spir_v_capability)New SPIR-V Capability

- [`ComputeDerivativeGroupQuadsNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ComputeDerivativeGroupQuadsKHR)
- [`ComputeDerivativeGroupLinearNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ComputeDerivativeGroupLinearKHR)

## [](#_issues)Issues

(1) Should we specify that the groups of four shader invocations used for derivatives in a compute shader are the same groups of four invocations that form a “quad” in shader subgroups?

**RESOLVED**: Yes.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2018-07-19 (Pat Brown)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_compute_shader_derivatives)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0