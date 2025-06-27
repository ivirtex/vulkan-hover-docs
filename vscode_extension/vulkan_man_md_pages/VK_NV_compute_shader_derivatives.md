# VK_NV_compute_shader_derivatives(3) Manual Page

## Name

VK_NV_compute_shader_derivatives - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

202

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_compute_shader_derivatives](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_compute_shader_derivatives.html)

## <a href="#_contact" class="anchor"></a>Contact

- Pat Brown <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_compute_shader_derivatives%5D%20@nvpbrown%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_compute_shader_derivatives%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>nvpbrown</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-07-19

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_compute_shader_derivatives`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_compute_shader_derivatives.txt)

**Contributors**  
- Pat Brown, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds Vulkan support for the
[`SPV_NV_compute_shader_derivatives`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_compute_shader_derivatives.html)
SPIR-V extension.

The SPIR-V extension provides two new execution modes, both of which
allow compute shaders to use built-ins that evaluate compute derivatives
explicitly or implicitly. Derivatives will be computed via differencing
over a 2x2 group of shader invocations. The `DerivativeGroupQuadsNV`
execution mode assembles shader invocations into 2x2 groups, where each
group has x and y coordinates of the local invocation ID of the form
(2m+{0,1}, 2n+{0,1}). The `DerivativeGroupLinearNV` execution mode
assembles shader invocations into 2x2 groups, where each group has local
invocation index values of the form 4m+{0,1,2,3}.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceComputeShaderDerivativesFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceComputeShaderDerivativesFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_COMPUTE_SHADER_DERIVATIVES_EXTENSION_NAME`

- `VK_NV_COMPUTE_SHADER_DERIVATIVES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV`

## <a href="#_new_spir_v_capability" class="anchor"></a>New SPIR-V Capability

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ComputeDerivativeGroupQuadsNV"
  target="_blank"
  rel="noopener"><code>ComputeDerivativeGroupQuadsNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ComputeDerivativeGroupLinearNV"
  target="_blank"
  rel="noopener"><code>ComputeDerivativeGroupLinearNV</code></a>

## <a href="#_issues" class="anchor"></a>Issues

\(1\) Should we specify that the groups of four shader invocations used
for derivatives in a compute shader are the same groups of four
invocations that form a “quad” in shader subgroups?

**RESOLVED**: Yes.

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-07-19 (Pat Brown)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_compute_shader_derivatives"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
