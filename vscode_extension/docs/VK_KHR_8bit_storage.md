# VK\_KHR\_8bit\_storage(3) Manual Page

## Name

VK\_KHR\_8bit\_storage - device extension



## [](#_registered_extension_number)Registered Extension Number

178

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK\_KHR\_storage\_buffer\_storage\_class](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_storage_buffer_storage_class.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_8bit\_storage](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_8bit_storage.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Alexander Galazin [\[GitHub\]alegal-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_8bit_storage%5D%20%40alegal-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_8bit_storage%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-02-05

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_shader_16bit_storage`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_shader_16bit_storage.txt)

**IP Status**

No known IP claims.

**Contributors**

- Alexander Galazin, Arm

## [](#_description)Description

The `VK_KHR_8bit_storage` extension allows use of 8-bit types in uniform and storage buffers, and push constant blocks. This extension introduces several new optional features which map to SPIR-V capabilities and allow access to 8-bit data in `Block`-decorated objects in the `Uniform` and the `StorageBuffer` storage classes, and objects in the `PushConstant` storage class.

The `StorageBuffer8BitAccess` capability **must** be supported by all implementations of this extension. The other capabilities are optional.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

Vulkan APIs in this extension are included in core Vulkan 1.2, with the KHR suffix omitted. However, if Vulkan 1.2 is supported and this extension is not, the `StorageBuffer8BitAccess` capability is optional. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

If Vulkan 1.4 is supported, support for the `storageBuffer8BitAccess` capability is required.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevice8BitStorageFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice8BitStorageFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_8BIT_STORAGE_EXTENSION_NAME`
- `VK_KHR_8BIT_STORAGE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`StorageBuffer8BitAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-StorageBuffer8BitAccess)
- [`UniformAndStorageBuffer8BitAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-UniformAndStorageBuffer8BitAccess)
- [`StoragePushConstant8`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-StoragePushConstant8)

## [](#_version_history)Version History

- Revision 1, 2018-02-05 (Alexander Galazin)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_8bit_storage).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0