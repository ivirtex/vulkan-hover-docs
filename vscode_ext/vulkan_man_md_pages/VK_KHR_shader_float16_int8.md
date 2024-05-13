# VK_KHR_shader_float16_int8(3) Manual Page

## Name

VK_KHR_shader_float16_int8 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

83

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Alexander Galazin <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_float16_int8%5D%20@alegal-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_float16_int8%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>alegal-arm</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-03-07

**Interactions and External Dependencies**  
- This extension interacts with
  [`VK_KHR_8bit_storage`](VK_KHR_8bit_storage.html)

- This extension interacts with
  [`VK_KHR_16bit_storage`](VK_KHR_16bit_storage.html)

- This extension interacts with
  [`VK_KHR_shader_float_controls`](VK_KHR_shader_float_controls.html)

- This extension provides API support for
  [`GL_EXT_shader_explicit_arithmetic_types`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_shader_explicit_arithmetic_types.txt)

**IP Status**  
No known IP claims.

**Contributors**  
- Alexander Galazin, Arm

- Jan-Harald Fredriksen, Arm

- Jeff Bolz, NVIDIA

- Graeme Leese, Broadcom

- Daniel Rakos, AMD

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_shader_float16_int8` extension allows use of 16-bit
floating-point types and 8-bit integer types in shaders for arithmetic
operations.

It introduces two new optional features `shaderFloat16` and `shaderInt8`
which directly map to the `Float16` and the `Int8` SPIR-V capabilities.
The `VK_KHR_shader_float16_int8` extension also specifies precision
requirements for half-precision floating-point SPIR-V operations. This
extension does not enable use of 8-bit integer types or 16-bit
floating-point types in any <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces"
target="_blank" rel="noopener">shader input and output interfaces</a>
and therefore does not supersede the
[`VK_KHR_8bit_storage`](VK_KHR_8bit_storage.html) or
[`VK_KHR_16bit_storage`](VK_KHR_16bit_storage.html) extensions.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. However, if Vulkan 1.2 is supported and this
extension is not, both the `shaderFloat16` and `shaderInt8` capabilities
are optional. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFloat16Int8FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFloat16Int8FeaturesKHR.html)

  - [VkPhysicalDeviceShaderFloat16Int8FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderFloat16Int8FeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_FLOAT16_INT8_EXTENSION_NAME`

- `VK_KHR_SHADER_FLOAT16_INT8_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT16_INT8_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-03-07 (Alexander Galazin)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_float16_int8"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
