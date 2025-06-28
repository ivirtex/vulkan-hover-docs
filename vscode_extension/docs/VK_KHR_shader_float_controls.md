# VK\_KHR\_shader\_float\_controls(3) Manual Page

## Name

VK\_KHR\_shader\_float\_controls - device extension



## [](#_registered_extension_number)Registered Extension Number

198

## [](#_revision)Revision

4

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_float\_controls](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_float_controls.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Alexander Galazin [\[GitHub\]alegal-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_float_controls%5D%20%40alegal-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_float_controls%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-09-11

**IP Status**

No known IP claims.

**Contributors**

- Alexander Galazin, Arm
- Jan-Harald Fredriksen, Arm
- Jeff Bolz, NVIDIA
- Graeme Leese, Broadcom
- Daniel Rakos, AMD

## [](#_description)Description

The `VK_KHR_shader_float_controls` extension enables efficient use of floating-point computations through the ability to query and override the implementation’s default behavior for rounding modes, denormals, signed zero, and infinity.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceFloatControlsPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFloatControlsPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkShaderFloatControlsIndependenceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderFloatControlsIndependenceKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_FLOAT_CONTROLS_EXTENSION_NAME`
- `VK_KHR_SHADER_FLOAT_CONTROLS_SPEC_VERSION`
- Extending [VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderFloatControlsIndependence.html):
  
  - `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR`
  - `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR`
  - `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`DenormPreserve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-DenormPreserve)
- [`DenormFlushToZero`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-DenormFlushToZero)
- [`SignedZeroInfNanPreserve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-SignedZeroInfNanPreserve)
- [`RoundingModeRTE`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RoundingModeRTE)
- [`RoundingModeRTZ`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RoundingModeRTZ)

## [](#_issues)Issues

1\) Which instructions must flush denorms?

**RESOLVED**: Only floating-point conversion, floating-point arithmetic, floating-point relational (except `OpIsNaN`, `OpIsInf`), and floating-point GLSL.std.450 extended instructions must flush denormals.

2\) What is the denorm behavior for intermediate results?

**RESOLVED**: When a SPIR-V instruction is implemented as a sequence of other instructions:

- in the `DenormFlushToZero` execution mode, the intermediate instructions may flush denormals, the final result of the sequence **must** not be denormal.
- in the `DenormPreserve` execution mode, denormals must be preserved throughout the whole sequence.

3\) Do denorm and rounding mode controls apply to `OpSpecConstantOp`?

**RESOLVED**: Yes, except when the opcode is `OpQuantizeToF16`.

4\) The SPIR-V specification says that `OpConvertFToU` and `OpConvertFToS` unconditionally round towards zero. Do the rounding mode controls specified through the execution modes apply to them?

**RESOLVED**: No, these instructions unconditionally round towards zero.

5\) Do any of the “Pack” GLSL.std.450 instructions count as conversion instructions and have the rounding mode applied?

**RESOLVED**: No, only instructions listed in “section 3.32.11. Conversion Instructions” of the SPIR-V specification count as conversion instructions.

6\) When using inf/nan-ignore mode, what is expected of `OpIsNan` and `OpIsInf`?

**RESOLVED**: These instructions must always accurately detect inf/nan if it is passed to them.

## [](#VK_KHR_shader_controls_v4_incompatibility)Version 4 API Incompatibility

The original versions of `VK_KHR_shader_float_controls` shipped with booleans named “separateDenormSettings” and “separateRoundingModeSettings”, which at first glance could have indicated “they can all be set independently, or not”. However the spec language as written indicated that the 32-bit value could always be set independently, and only the 16- and 64-bit controls needed to be the same if these values were `VK_FALSE`.

As a result of this slight disparity, and lack of test coverage for this facet of the extension, we ended up with two different behaviors in the wild, where some implementations worked as written, and others worked based on the naming. As these are hard limits in hardware with reasons for exposure as written, it was not possible to standardize on a single way to make this work within the existing API.

No known users of this part of the extension exist in the wild, and as such the Vulkan WG took the unusual step of retroactively changing the once boolean value into a tri-state enum, breaking source compatibility. This was however done in such a way as to retain ABI compatibility, in case any code using this did exist; with the numerical values 0 and 1 retaining their original specified meaning, and a new value signifying the additional “all need to be set together” state. If any applications exist today, compiled binaries will continue to work as written in most cases, but will need changes before the code can be recompiled.

## [](#_version_history)Version History

- Revision 4, 2019-06-18 (Tobias Hector)
  
  - Modified settings restrictions, see [Version 4 API incompatibility](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_controls_v4_incompatibility.html)
- Revision 3, 2018-09-11 (Alexander Galazin)
  
  - Minor restructuring
- Revision 2, 2018-04-17 (Alexander Galazin)
  
  - Added issues and resolutions
- Revision 1, 2018-04-11 (Alexander Galazin)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_float_controls)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0