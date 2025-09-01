# VK\_KHR\_fragment\_shader\_barycentric(3) Manual Page

## Name

VK\_KHR\_fragment\_shader\_barycentric - device extension



## [](#_registered_extension_number)Registered Extension Number

323

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_fragment\_shader\_barycentric](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_fragment_shader_barycentric.html)

## [](#_contact)Contact

- Stu Smith

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_fragment\_shader\_barycentric](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_fragment_shader_barycentric.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-03-10

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_fragment_shader_barycentric`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_fragment_shader_barycentric.txt)

**Contributors**

- Stu Smith, AMD
- Tobias Hector, AMD
- Graeme Leese, Broadcom
- Jan-Harald Fredriksen, Arm
- Slawek Grajewski, Intel
- Pat Brown, NVIDIA
- Hans-Kristian Arntzen, Valve
- Contributors to the VK\_NV\_fragment\_shader\_barycentric specification

## [](#_description)Description

This extension is based on the `VK_NV_fragment_shader_barycentric` extension, and adds support for the following SPIR-V extension in Vulkan:

- [`SPV_KHR_fragment_shader_barycentric`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_fragment_shader_barycentric.html)

The extension provides access to three additional fragment shader variable decorations in SPIR-V:

- `PerVertexKHR`, which indicates that a fragment shader input will not have interpolated values, but instead must be accessed with an extra array index that identifies one of the vertices of the primitive producing the fragment
- `BaryCoordKHR`, which indicates that the variable is a three-component floating-point vector holding barycentric weights for the fragment produced using perspective interpolation
- `BaryCoordNoPerspKHR`, which indicates that the variable is a three-component floating-point vector holding barycentric weights for the fragment produced using linear interpolation

When using GLSL source-based shader languages, the following variables from `GL_EXT_fragment_shader_barycentric` map to these SPIR-V built-in decorations:

- `in vec3 gl_BaryCoordEXT;` → `BaryCoordKHR`
- `in vec3 gl_BaryCoordNoPerspEXT;` → `BaryCoordNoPerspKHR`

GLSL variables declared using the `pervertexEXT` GLSL qualifier are expected to be decorated with `PerVertexKHR` in SPIR-V.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_FRAGMENT_SHADER_BARYCENTRIC_EXTENSION_NAME`
- `VK_KHR_FRAGMENT_SHADER_BARYCENTRIC_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_PROPERTIES_KHR`

## [](#_new_built_in_variables)New Built-In Variables

- [`BaryCoordKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-barycoordkhr)
- [`BaryCoordNoPerspKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-barycoordnoperspkhr)

## [](#_new_spir_v_decorations)New SPIR-V Decorations

- [`PerVertexKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-interpolation-decorations-pervertexkhr)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`FragmentBarycentricKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-FragmentBarycentricKHR)

## [](#_issues)Issues

1\) What are the interactions with MSAA and how are `BaryCoordKHR` and `BaryCoordNoPerspKHR` interpolated?

**RESOLVED**: The inputs decorated with `BaryCoordKHR` or `BaryCoordNoPerspKHR` **may** also be decorated with the `Centroid` or `Sample` qualifiers to specify interpolation, like any other fragment shader input. If the [`shaderSampleRateInterpolationFunctions`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSampleRateInterpolationFunctions) feature is enabled, the extended instructions InterpolateAtCentroid, InterpolateAtOffset, and InterpolateAtSample from the GLSL.std.450 **may** also be used with inputs decorated with `BaryCoordKHR` or `BaryCoordNoPerspKHR`.

## [](#_version_history)Version History

- Revision 1, 2022-03-10 (Stu Smith)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_fragment_shader_barycentric).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0