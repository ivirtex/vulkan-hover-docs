# VK_KHR_fragment_shader_barycentric(3) Manual Page

## Name

VK_KHR_fragment_shader_barycentric - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

323

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_fragment_shader_barycentric](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_fragment_shader_barycentric.html)

## <a href="#_contact" class="anchor"></a>Contact

- Stu Smith

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_fragment_shader_barycentric](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_fragment_shader_barycentric.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-03-10

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_EXT_fragment_shader_barycentric`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_fragment_shader_barycentric.txt)

**Contributors**  
- Stu Smith, AMD

- Tobias Hector, AMD

- Graeme Leese, Broadcom

- Jan-Harald Fredriksen, Arm

- Slawek Grajewski, Intel

- Pat Brown, NVIDIA

- Hans-Kristian Arntzen, Valve

- Contributors to the VK_NV_fragment_shader_barycentric specification

## <a href="#_description" class="anchor"></a>Description

This extension is based on the
[`VK_NV_fragment_shader_barycentric`](VK_NV_fragment_shader_barycentric.html)
extension, and adds support for the following SPIR-V extension in
Vulkan:

- [`SPV_KHR_fragment_shader_barycentric`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_fragment_shader_barycentric.html)

The extension provides access to three additional fragment shader
variable decorations in SPIR-V:

- `PerVertexKHR`, which indicates that a fragment shader input will not
  have interpolated values, but instead must be accessed with an extra
  array index that identifies one of the vertices of the primitive
  producing the fragment

- `BaryCoordKHR`, which indicates that the variable is a three-component
  floating-point vector holding barycentric weights for the fragment
  produced using perspective interpolation

- `BaryCoordNoPerspKHR`, which indicates that the variable is a
  three-component floating-point vector holding barycentric weights for
  the fragment produced using linear interpolation

When using GLSL source-based shader languages, the following variables
from `GL_EXT_fragment_shader_barycentric` map to these SPIR-V built-in
decorations:

- `in vec3 gl_BaryCoordEXT;` → `BaryCoordKHR`

- `in vec3 gl_BaryCoordNoPerspEXT;` → `BaryCoordNoPerspKHR`

GLSL variables declared using the `pervertexEXT` GLSL qualifier are
expected to be decorated with `PerVertexKHR` in SPIR-V.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_FRAGMENT_SHADER_BARYCENTRIC_EXTENSION_NAME`

- `VK_KHR_FRAGMENT_SHADER_BARYCENTRIC_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_PROPERTIES_KHR`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-barycoordkhr"
  target="_blank" rel="noopener"><code>BaryCoordKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-barycoordnoperspkhr"
  target="_blank" rel="noopener"><code>BaryCoordNoPerspKHR</code></a>

## <a href="#_new_spir_v_decorations" class="anchor"></a>New SPIR-V Decorations

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-interpolation-decorations-pervertexkhr"
  target="_blank" rel="noopener"><code>PerVertexKHR</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FragmentBarycentricKHR"
  target="_blank" rel="noopener"><code>FragmentBarycentricKHR</code></a>

## <a href="#_issues" class="anchor"></a>Issues

1\) What are the interactions with MSAA and how are `BaryCoordKHR` and
`BaryCoordNoPerspKHR` interpolated?

**RESOLVED**: The inputs decorated with `BaryCoordKHR` or
`BaryCoordNoPerspKHR` **may** also be decorated with the `Centroid` or
`Sample` qualifiers to specify interpolation, like any other fragment
shader input. If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSampleRateInterpolationFunctions"
target="_blank"
rel="noopener"><code>shaderSampleRateInterpolationFunctions</code></a>
is enabled, the extended instructions InterpolateAtCentroid,
InterpolateAtOffset, and InterpolateAtSample from the GLSL.std.450
**may** also be used with inputs decorated with `BaryCoordKHR` or
`BaryCoordNoPerspKHR`.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-03-10 (Stu Smith)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_fragment_shader_barycentric"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
