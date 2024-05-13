# VK_KHR_shader_terminate_invocation(3) Manual Page

## Name

VK_KHR_shader_terminate_invocation - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

216

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_terminate_invocation](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_terminate_invocation.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_terminate_invocation%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_terminate_invocation%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>critsec</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-08-11

**IP Status**  
No known IP claims.

**Contributors**  
- Alan Baker, Google

- Jeff Bolz, NVIDIA

- Jesse Hall, Google

- Ralph Potter, Samsung

- Tom Olson, Arm

## <a href="#_description" class="anchor"></a>Description

This extension adds Vulkan support for the
[`SPV_KHR_terminate_invocation`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_terminate_invocation.html)
SPIR-V extension. That SPIR-V extension provides a new instruction,
`OpTerminateInvocation`, which causes a shader invocation to immediately
terminate and sets the coverage of shaded samples to `0`; only
previously executed instructions will have observable effects. The
`OpTerminateInvocation` instruction, along with the
`OpDemoteToHelperInvocation` instruction from the
[`VK_EXT_shader_demote_to_helper_invocation`](VK_EXT_shader_demote_to_helper_invocation.html)
extension, together replace the `OpKill` instruction, which could behave
like either of these instructions. `OpTerminateInvocation` provides the
behavior required by the GLSL `discard` statement, and should be used
when available by GLSL compilers and applications that need the GLSL
`discard` behavior.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderTerminateInvocationFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderTerminateInvocationFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_TERMINATE_INVOCATION_EXTENSION_NAME`

- `VK_KHR_SHADER_TERMINATE_INVOCATION_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TERMINATE_INVOCATION_FEATURES_KHR`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
KHR suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-08-11 (Jesse Hall)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_terminate_invocation"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
