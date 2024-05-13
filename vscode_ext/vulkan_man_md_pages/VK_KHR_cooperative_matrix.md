# VK_KHR_cooperative_matrix(3) Manual Page

## Name

VK_KHR_cooperative_matrix - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

507

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_cooperative_matrix](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_cooperative_matrix.html)

## <a href="#_contact" class="anchor"></a>Contact

- Kevin Petit <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_cooperative_matrix%5D%20@kpet%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_cooperative_matrix%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kpet</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_cooperative_matrix](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_cooperative_matrix.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-05-03

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_KHR_cooperative_matrix`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/khr/GLSL_KHR_cooperative_matrix.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

- Markus Tavenrath, NVIDIA

- Daniel Koch, NVIDIA

- Kevin Petit, Arm Ltd.

- Boris Zanin, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds support for using cooperative matrix types in
SPIR-V. Cooperative matrix types are medium-sized matrices that are
primarily supported in compute shaders, where the storage for the matrix
is spread across all invocations in some scope (usually a subgroup) and
those invocations cooperate to efficiently perform matrix multiplies.

Cooperative matrix types are defined by the
[`SPV_KHR_cooperative_matrix`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_cooperative_matrix.html)
SPIR-V extension and can be used with the
[`GLSL_KHR_cooperative_matrix`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/khr/GLSL_KHR_cooperative_matrix.txt)
GLSL extension.

This extension includes support for enumerating the matrix types and
dimensions that are supported by the implementation.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCooperativeMatrixFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeKHR.html)

- [VkScopeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScopeKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_COOPERATIVE_MATRIX_EXTENSION_NAME`

- `VK_KHR_COOPERATIVE_MATRIX_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixKHR"
  target="_blank" rel="noopener">CooperativeMatrixKHR</a>

## <a href="#_issues" class="anchor"></a>Issues

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2023-05-03 (Kevin Petit)

  - First KHR revision

- Revision 1, 2019-02-05 (Jeff Bolz)

  - NVIDIA vendor extension

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_cooperative_matrix"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
