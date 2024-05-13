# VK_KHR_shader_float_controls(3) Manual Page

## Name

VK_KHR_shader_float_controls - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

198

## <a href="#_revision" class="anchor"></a>Revision

4

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_float_controls](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_float_controls.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Alexander Galazin <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_float_controls%5D%20@alegal-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_float_controls%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>alegal-arm</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_shader_float_controls` extension enables efficient use of
floating-point computations through the ability to query and override
the implementation’s default behavior for rounding modes, denormals,
signed zero, and infinity.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceFloatControlsPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFloatControlsPropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkShaderFloatControlsIndependenceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependenceKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_FLOAT_CONTROLS_EXTENSION_NAME`

- `VK_KHR_SHADER_FLOAT_CONTROLS_SPEC_VERSION`

- Extending
  [VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html):

  - `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR`

  - `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR`

  - `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-DenormPreserve"
  target="_blank" rel="noopener"><code>DenormPreserve</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-DenormFlushToZero"
  target="_blank" rel="noopener"><code>DenormFlushToZero</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-SignedZeroInfNanPreserve"
  target="_blank" rel="noopener"><code>SignedZeroInfNanPreserve</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RoundingModeRTE"
  target="_blank" rel="noopener"><code>RoundingModeRTE</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RoundingModeRTZ"
  target="_blank" rel="noopener"><code>RoundingModeRTZ</code></a>

## <a href="#_issues" class="anchor"></a>Issues

1\) Which instructions must flush denorms?

**RESOLVED**: Only floating-point conversion, floating-point arithmetic,
floating-point relational (except `OpIsNaN`, `OpIsInf`), and
floating-point GLSL.std.450 extended instructions must flush denormals.

2\) What is the denorm behavior for intermediate results?

**RESOLVED**: When a SPIR-V instruction is implemented as a sequence of
other instructions:

- in the `DenormFlushToZero` execution mode, the intermediate
  instructions may flush denormals, the final result of the sequence
  **must** not be denormal.

- in the `DenormPreserve` execution mode, denormals must be preserved
  throughout the whole sequence.

3\) Do denorm and rounding mode controls apply to `OpSpecConstantOp`?

**RESOLVED**: Yes, except when the opcode is `OpQuantizeToF16`.

4\) The SPIR-V specification says that `OpConvertFToU` and
`OpConvertFToS` unconditionally round towards zero. Do the rounding mode
controls specified through the execution modes apply to them?

**RESOLVED**: No, these instructions unconditionally round towards zero.

5\) Do any of the “Pack” GLSL.std.450 instructions count as conversion
instructions and have the rounding mode applied?

**RESOLVED**: No, only instructions listed in “section 3.32.11.
Conversion Instructions” of the SPIR-V specification count as conversion
instructions.

6\) When using inf/nan-ignore mode, what is expected of `OpIsNan` and
`OpIsInf`?

**RESOLVED**: These instructions must always accurately detect inf/nan
if it is passed to them.

## <a href="#VK_KHR_shader_controls_v4_incompatibility" class="anchor"></a>Version 4 API Incompatibility

The original versions of `VK_KHR_shader_float_controls` shipped with
booleans named “separateDenormSettings” and
“separateRoundingModeSettings”, which at first glance could have
indicated “they can all be set independently, or not”. However the spec
language as written indicated that the 32-bit value could always be set
independently, and only the 16- and 64-bit controls needed to be the
same if these values were `VK_FALSE`.

As a result of this slight disparity, and lack of test coverage for this
facet of the extension, we ended up with two different behaviors in the
wild, where some implementations worked as written, and others worked
based on the naming. As these are hard limits in hardware with reasons
for exposure as written, it was not possible to standardize on a single
way to make this work within the existing API.

No known users of this part of the extension exist in the wild, and as
such the Vulkan WG took the unusual step of retroactively changing the
once boolean value into a tri-state enum, breaking source compatibility.
This was however done in such a way as to retain ABI compatibility, in
case any code using this did exist; with the numerical values 0 and 1
retaining their original specified meaning, and a new value signifying
the additional “all need to be set together” state. If any applications
exist today, compiled binaries will continue to work as written in most
cases, but will need changes before the code can be recompiled.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 4, 2019-06-18 (Tobias Hector)

  - Modified settings restrictions, see
    <a href="VK_KHR_shader_controls_v4_incompatibility.html" target="_blank"
    rel="noopener">Version 4 API incompatibility</a>

- Revision 3, 2018-09-11 (Alexander Galazin)

  - Minor restructuring

- Revision 2, 2018-04-17 (Alexander Galazin)

  - Added issues and resolutions

- Revision 1, 2018-04-11 (Alexander Galazin)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_float_controls"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
