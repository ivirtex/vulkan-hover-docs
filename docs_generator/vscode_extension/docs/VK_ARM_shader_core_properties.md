# VK\_ARM\_shader\_core\_properties(3) Manual Page

## Name

VK\_ARM\_shader\_core\_properties - device extension



## [](#_registered_extension_number)Registered Extension Number

416

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_shader_core_properties%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_shader_core_properties%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-02-07

**IP Status**

No known IP claims.

**Contributors**

- Jan-Harald Fredriksen, Arm Ltd.

## [](#_description)Description

This extension provides the ability to determine device-specific performance properties of Arm GPUs.

It exposes properties for the number of texel, pixel, and fused multiply-add operations per clock per shader core. This can be used in combination with the `VK_ARM_shader_core_builtins` extension that provides the ability to query the number of shader cores on the physical device.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderCorePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderCorePropertiesARM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ARM_SHADER_CORE_PROPERTIES_EXTENSION_NAME`
- `VK_ARM_SHADER_CORE_PROPERTIES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_ARM`

## [](#_version_history)Version History

- Revision 1, 2023-02-07 (Jan-Harald Fredriksen)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ARM_shader_core_properties)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0