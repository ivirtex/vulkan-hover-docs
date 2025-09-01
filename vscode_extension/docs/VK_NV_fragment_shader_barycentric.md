# VK\_NV\_fragment\_shader\_barycentric(3) Manual Page

## Name

VK\_NV\_fragment\_shader\_barycentric - device extension



## [](#_registered_extension_number)Registered Extension Number

204

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_fragment\_shader\_barycentric](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_fragment_shader_barycentric.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_fragment\_shader\_barycentric](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shader_barycentric.html) extension

## [](#_contact)Contact

- Pat Brown [\[GitHub\]nvpbrown](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_fragment_shader_barycentric%5D%20%40nvpbrown%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_fragment_shader_barycentric%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-08-03

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_fragment_shader_barycentric`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_fragment_shader_barycentric.txt)

**Contributors**

- Pat Brown, NVIDIA
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- [`SPV_NV_fragment_shader_barycentric`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_fragment_shader_barycentric.html)

The extension provides access to three additional fragment shader variable decorations in SPIR-V:

- `PerVertexNV`, which indicates that a fragment shader input will not have interpolated values, but instead must be accessed with an extra array index that identifies one of the vertices of the primitive producing the fragment
- `BaryCoordNV`, which indicates that the variable is a three-component floating-point vector holding barycentric weights for the fragment produced using perspective interpolation
- `BaryCoordNoPerspNV`, which indicates that the variable is a three-component floating-point vector holding barycentric weights for the fragment produced using linear interpolation

When using GLSL source-based shader languages, the following variables from `GL_NV_fragment_shader_barycentric` maps to these SPIR-V built-in decorations:

- `in vec3 gl_BaryCoordNV;` → `BaryCoordNV`
- `in vec3 gl_BaryCoordNoPerspNV;` → `BaryCoordNoPerspNV`

GLSL variables declared using the `__pervertexNV` GLSL qualifier are expected to be decorated with `PerVertexNV` in SPIR-V.

## [](#_promotion_to_vk_khr_fragment_shader_barycentric)Promotion to `VK_KHR_fragment_shader_barycentric`

All functionality in this extension is included in `VK_KHR_fragment_shader_barycentric`, with the suffix changed to KHR.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFragmentShaderBarycentricFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentShaderBarycentricFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_FRAGMENT_SHADER_BARYCENTRIC_EXTENSION_NAME`
- `VK_NV_FRAGMENT_SHADER_BARYCENTRIC_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_NV`

## [](#_new_built_in_variables)New Built-In Variables

- [`BaryCoordNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-barycoordkhr)
- [`BaryCoordNoPerspNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-barycoordnoperspkhr)

## [](#_new_spir_v_decorations)New SPIR-V Decorations

- [`PerVertexNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-interpolation-decorations-pervertexkhr)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`FragmentBarycentricNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-FragmentBarycentricKHR)

## [](#_issues)Issues

(1) The AMD\_shader\_explicit\_vertex\_parameter extension provides similar functionality. Why write a new extension, and how is this extension different?

**RESOLVED**: For the purposes of Vulkan/SPIR-V, we chose to implement a separate extension due to several functional differences.

First, the hardware supporting this extension can provide a three-component barycentric weight vector for variables decorated with `BaryCoordNV`, while variables decorated with `BaryCoordSmoothAMD` provide only two components. In some cases, it may be more efficient to explicitly interpolate an attribute via:

```
float value = (baryCoordNV.x * v[0].attrib +
               baryCoordNV.y * v[1].attrib +
               baryCoordNV.z * v[2].attrib);
```

instead of

```
float value = (baryCoordSmoothAMD.x * (v[0].attrib - v[2].attrib) +
               baryCoordSmoothAMD.y * (v[1].attrib - v[2].attrib) +
               v[2].attrib);
```

Additionally, the semantics of the decoration `BaryCoordPullModelAMD` do not appear to map to anything supported by the initial hardware implementation of this extension.

This extension provides a smaller number of decorations than the AMD extension, as we expect that shaders could derive variables decorated with things like `BaryCoordNoPerspCentroidAMD` with explicit attribute interpolation instructions. One other relevant difference is that explicit per-vertex attribute access using this extension does not require a constant vertex number.

(2) Why do the built-in SPIR-V decorations for this extension include two separate built-ins `BaryCoordNV` and `BaryCoordNoPerspNV` when a “no perspective” variable could be decorated with `BaryCoordNV` and `NoPerspective`?

**RESOLVED**: The SPIR-V extension for this feature chose to mirror the behavior of the GLSL extension, which provides two built-in variables. Additionally, it is not clear that its a good idea (or even legal) to have two variables using the “same attribute”, but with different interpolation modifiers.

## [](#_version_history)Version History

- Revision 1, 2018-08-03 (Pat Brown)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_fragment_shader_barycentric).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0