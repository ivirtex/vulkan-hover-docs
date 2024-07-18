# VK_KHR_shader_quad_control(3) Manual Page

## Name

VK_KHR_shader_quad_control - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

236

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  
and  
[VK_KHR_vulkan_memory_model](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vulkan_memory_model.html)  
and  
[VK_KHR_shader_maximal_reconvergence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_maximal_reconvergence.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_quad_control](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_quad_control.html)

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_quad_control%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_quad_control%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_shader_quad_control](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_quad_control.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-11-01

**IP Status**  
No known IP claims.

**Contributors**  
- Tobias Hector, AMD

- Bill Licea-Kane, Qualcomm

- Graeme Leese, Broadcom

- Jan-Harald Fredriksen, Arm

- Nicolai Hähnle, AMD

- Jeff Bolz, NVidia

- Alan Baker, Google

- Hans-Kristian Arntzen, Valve

## <a href="#_description" class="anchor"></a>Description

This extension adds new quad any/all operations, requires that
derivatives are well-defined in quad-uniform control flow, and adds the
ability to require helper invocations participate in group operations.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderQuadControlFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderQuadControlFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_QUAD_CONTROL_EXTENSION_NAME`

- `VK_KHR_SHADER_QUAD_CONTROL_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_QUAD_CONTROL_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-QuadControlKHR"
  target="_blank" rel="noopener">QuadControlKHR</a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-11-01 (Tobias Hector)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_quad_control"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
