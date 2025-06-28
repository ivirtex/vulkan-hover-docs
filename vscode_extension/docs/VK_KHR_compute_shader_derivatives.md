# VK\_KHR\_compute\_shader\_derivatives(3) Manual Page

## Name

VK\_KHR\_compute\_shader\_derivatives - device extension



## [](#_registered_extension_number)Registered Extension Number

512

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_compute\_shader\_derivatives](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_compute_shader_derivatives.html)

## [](#_contact)Contact

- Jean-Noe Morissette [\[GitHub\]MagicPoncho](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_compute_shader_derivatives%5D%20%40MagicPoncho%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_compute_shader_derivatives%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_compute\_shader\_derivatives](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_compute_shader_derivatives.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-06-26

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension requires [`SPV_KHR_compute_shader_derivatives`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_compute_shader_derivatives.html)
- This extension provides API support for [`GL_KHR_compute_shader_derivatives`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/khr/GLSL_KHR_compute_shader_derivatives.txt)

**Contributors**

- Jean-Noe Morissette, Epic Games
- Daniel Koch, NVIDIA
- Pat Brown, NVIDIA
- Stu Smith, AMD
- Jan-Harald Fredriksen, Arm
- Tobias Hector, AMD
- Ralph Potter, Samsung
- Pan Gao, Huawei
- Samuel (Sheng-Wen) Huang, MediaTek
- Graeme Leese, Broadcom
- Hans-Kristian Arntzen, Valve
- Matthew Netsch, Qualcomm

## [](#_description)Description

This extension adds Vulkan support for the [`SPV_KHR_compute_shader_derivatives`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_compute_shader_derivatives.html) SPIR-V extension.

The SPIR-V extension provides two new execution modes, both of which allow execution models with defined workgroups to use built-ins that evaluate derivatives explicitly or implicitly. Derivatives will be computed via differencing over a 2x2 group of shader invocations. The `DerivativeGroupQuadsKHR` execution mode assembles shader invocations into 2x2 groups, where each group has x and y coordinates of the local invocation ID of the form (2m+{0,1}, 2n+{0,1}). The `DerivativeGroupLinearKHR` execution mode assembles shader invocations into 2x2 groups, where each group has local invocation index values of the form 4m+{0,1,2,3}.

The new execution modes are supported in compute shaders and optionally (see [meshAndTaskShaderDerivatives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-meshAndTaskShaderDerivatives)) in mesh and task shaders.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceComputeShaderDerivativesFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceComputeShaderDerivativesFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_COMPUTE_SHADER_DERIVATIVES_EXTENSION_NAME`
- `VK_KHR_COMPUTE_SHADER_DERIVATIVES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_PROPERTIES_KHR`

## [](#_new_spir_v_capability)New SPIR-V Capability

- [`ComputeDerivativeGroupQuadsKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ComputeDerivativeGroupQuadsKHR)
- [`ComputeDerivativeGroupLinearKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ComputeDerivativeGroupLinearKHR)

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2023-02-27 (Jean-Noe Morissette)
  
  - Initial draft
  - Add properties and clarify mesh and task support (Daniel Koch)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_compute_shader_derivatives)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0