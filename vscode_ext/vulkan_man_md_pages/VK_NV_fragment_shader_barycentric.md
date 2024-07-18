# VK_NV_fragment_shader_barycentric(3) Manual Page

## Name

VK_NV_fragment_shader_barycentric - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

204

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_fragment_shader_barycentric](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_fragment_shader_barycentric.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to
  [VK_KHR_fragment_shader_barycentric](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shader_barycentric.html)
  extension

## <a href="#_contact" class="anchor"></a>Contact

- Pat Brown <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_fragment_shader_barycentric%5D%20@nvpbrown%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_fragment_shader_barycentric%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>nvpbrown</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-08-03

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_fragment_shader_barycentric`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_fragment_shader_barycentric.txt)

**Contributors**  
- Pat Brown, NVIDIA

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the following SPIR-V extension in
Vulkan:

- [`SPV_NV_fragment_shader_barycentric`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_fragment_shader_barycentric.html)

The extension provides access to three additional fragment shader
variable decorations in SPIR-V:

- `PerVertexNV`, which indicates that a fragment shader input will not
  have interpolated values, but instead must be accessed with an extra
  array index that identifies one of the vertices of the primitive
  producing the fragment

- `BaryCoordNV`, which indicates that the variable is a three-component
  floating-point vector holding barycentric weights for the fragment
  produced using perspective interpolation

- `BaryCoordNoPerspNV`, which indicates that the variable is a
  three-component floating-point vector holding barycentric weights for
  the fragment produced using linear interpolation

When using GLSL source-based shader languages, the following variables
from `GL_NV_fragment_shader_barycentric` maps to these SPIR-V built-in
decorations:

- `in vec3 gl_BaryCoordNV;` → `BaryCoordNV`

- `in vec3 gl_BaryCoordNoPerspNV;` → `BaryCoordNoPerspNV`

GLSL variables declared using the `__pervertexNV` GLSL qualifier are
expected to be decorated with `PerVertexNV` in SPIR-V.

## <a href="#_promotion_to_vk_khr_fragment_shader_barycentric"
class="anchor"></a>Promotion to `VK_KHR_fragment_shader_barycentric`

All functionality in this extension is included in
[`VK_KHR_fragment_shader_barycentric`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shader_barycentric.html),
with the suffix changed to KHR.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFragmentShaderBarycentricFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShaderBarycentricFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_FRAGMENT_SHADER_BARYCENTRIC_EXTENSION_NAME`

- `VK_NV_FRAGMENT_SHADER_BARYCENTRIC_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_NV`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-barycoordkhr"
  target="_blank" rel="noopener"><code>BaryCoordNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-barycoordnoperspkhr"
  target="_blank" rel="noopener"><code>BaryCoordNoPerspNV</code></a>

## <a href="#_new_spir_v_decorations" class="anchor"></a>New SPIR-V Decorations

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-interpolation-decorations-pervertexkhr"
  target="_blank" rel="noopener"><code>PerVertexNV</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FragmentBarycentricKHR"
  target="_blank" rel="noopener"><code>FragmentBarycentricNV</code></a>

## <a href="#_issues" class="anchor"></a>Issues

\(1\) The AMD_shader_explicit_vertex_parameter extension provides
similar functionality. Why write a new extension, and how is this
extension different?

**RESOLVED**: For the purposes of Vulkan/SPIR-V, we chose to implement a
separate extension due to several functional differences.

First, the hardware supporting this extension can provide a
three-component barycentric weight vector for variables decorated with
`BaryCoordNV`, while variables decorated with `BaryCoordSmoothAMD`
provide only two components. In some cases, it may be more efficient to
explicitly interpolate an attribute via:

    float value = (baryCoordNV.x * v[0].attrib +
                   baryCoordNV.y * v[1].attrib +
                   baryCoordNV.z * v[2].attrib);

instead of

    float value = (baryCoordSmoothAMD.x * (v[0].attrib - v[2].attrib) +
                   baryCoordSmoothAMD.y * (v[1].attrib - v[2].attrib) +
                   v[2].attrib);

Additionally, the semantics of the decoration `BaryCoordPullModelAMD` do
not appear to map to anything supported by the initial hardware
implementation of this extension.

This extension provides a smaller number of decorations than the AMD
extension, as we expect that shaders could derive variables decorated
with things like `BaryCoordNoPerspCentroidAMD` with explicit attribute
interpolation instructions. One other relevant difference is that
explicit per-vertex attribute access using this extension does not
require a constant vertex number.

\(2\) Why do the built-in SPIR-V decorations for this extension include
two separate built-ins `BaryCoordNV` and `BaryCoordNoPerspNV` when a “no
perspective” variable could be decorated with `BaryCoordNV` and
`NoPerspective`?

**RESOLVED**: The SPIR-V extension for this feature chose to mirror the
behavior of the GLSL extension, which provides two built-in variables.
Additionally, it is not clear that its a good idea (or even legal) to
have two variables using the “same attribute”, but with different
interpolation modifiers.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-08-03 (Pat Brown)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_fragment_shader_barycentric"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
