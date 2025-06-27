# VK_INTEL_shader_integer_functions2(3) Manual Page

## Name

VK_INTEL_shader_integer_functions2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

210

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_INTEL_shader_integer_functions2](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/INTEL/SPV_INTEL_shader_integer_functions2.html)

## <a href="#_contact" class="anchor"></a>Contact

- Ian Romanick <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_INTEL_shader_integer_functions2%5D%20@ianromanick%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_INTEL_shader_integer_functions2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ianromanick</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-04-30

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_INTEL_shader_integer_functions2`](https://registry.khronos.org/OpenGL/extensions/INTEL/INTEL_shader_integer_functions2.txt).

**Contributors**  
- Ian Romanick, Intel

- Ben Ashbaugh, Intel

## <a href="#_description" class="anchor"></a>Description

This extension adds support for several new integer instructions in
SPIR-V for use in graphics shaders. Many of these instructions have
pre-existing counterparts in the Kernel environment.

The added integer functions are defined by the
[`SPV_INTEL_shader_integer_functions2`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/INTEL/SPV_INTEL_shader_integer_functions2.html)
SPIR-V extension and can be used with the
`GL_INTEL_shader_integer_functions2` GLSL extension.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_INTEL_SHADER_INTEGER_FUNCTIONS_2_EXTENSION_NAME`

- `VK_INTEL_SHADER_INTEGER_FUNCTIONS_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_FUNCTIONS_2_FEATURES_INTEL`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-IntegerFunctions2INTEL"
  target="_blank" rel="noopener"><code>IntegerFunctions2INTEL</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-04-30 (Ian Romanick)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_INTEL_shader_integer_functions2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
