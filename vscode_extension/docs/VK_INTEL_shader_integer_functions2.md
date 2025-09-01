# VK\_INTEL\_shader\_integer\_functions2(3) Manual Page

## Name

VK\_INTEL\_shader\_integer\_functions2 - device extension



## [](#_registered_extension_number)Registered Extension Number

210

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_INTEL\_shader\_integer\_functions2](https://github.khronos.org/SPIRV-Registry/extensions/INTEL/SPV_INTEL_shader_integer_functions2.html)

## [](#_contact)Contact

- Ian Romanick [\[GitHub\]ianromanick](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_INTEL_shader_integer_functions2%5D%20%40ianromanick%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_INTEL_shader_integer_functions2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-04-30

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_INTEL_shader_integer_functions2`](https://registry.khronos.org/OpenGL/extensions/INTEL/INTEL_shader_integer_functions2.txt).

**Contributors**

- Ian Romanick, Intel
- Ben Ashbaugh, Intel

## [](#_description)Description

This extension adds support for several new integer instructions in SPIR-V for use in graphics shaders. Many of these instructions have pre-existing counterparts in the Kernel environment.

The added integer functions are defined by the [`SPV_INTEL_shader_integer_functions2`](https://github.khronos.org/SPIRV-Registry/extensions/INTEL/SPV_INTEL_shader_integer_functions2.html) SPIR-V extension and can be used with the `GL_INTEL_shader_integer_functions2` GLSL extension.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_INTEL_SHADER_INTEGER_FUNCTIONS_2_EXTENSION_NAME`
- `VK_INTEL_SHADER_INTEGER_FUNCTIONS_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_FUNCTIONS_2_FEATURES_INTEL`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`IntegerFunctions2INTEL`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-IntegerFunctions2INTEL)

## [](#_version_history)Version History

- Revision 1, 2019-04-30 (Ian Romanick)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_INTEL_shader_integer_functions2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0