# VK\_KHR\_shader\_quad\_control(3) Manual Page

## Name

VK\_KHR\_shader\_quad\_control - device extension



## [](#_registered_extension_number)Registered Extension Number

236

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [Vulkan Version 1.1](#versions-1.1)  
         and  
         [VK\_KHR\_vulkan\_memory\_model](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vulkan_memory_model.html)  
     or  
     [Vulkan Version 1.2](#versions-1.2)  
and  
[VK\_KHR\_shader\_maximal\_reconvergence](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_maximal_reconvergence.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_quad\_control](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_quad_control.html)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_quad_control%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_quad_control%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_quad\_control](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_quad_control.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

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

## [](#_description)Description

This extension adds new quad any/all operations, requires that derivatives are well-defined in quad-uniform control flow, and adds the ability to require helper invocations participate in group operations.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderQuadControlFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderQuadControlFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_QUAD_CONTROL_EXTENSION_NAME`
- `VK_KHR_SHADER_QUAD_CONTROL_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_QUAD_CONTROL_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [QuadControlKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-QuadControlKHR)

## [](#_version_history)Version History

- Revision 1, 2023-11-01 (Tobias Hector)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_quad_control).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0