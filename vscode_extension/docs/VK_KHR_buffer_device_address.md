# VK\_KHR\_buffer\_device\_address(3) Manual Page

## Name

VK\_KHR\_buffer\_device\_address - device extension



## [](#_registered_extension_number)Registered Extension Number

258

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_physical\_storage\_buffer](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_physical_storage_buffer.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_buffer_device_address%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_buffer_device_address%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-06-24

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_buffer_reference`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference.txt) and [`GL_EXT_buffer_reference2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference2.txt) and [`GL_EXT_buffer_reference_uvec2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference_uvec2.txt)

**Contributors**

- Jeff Bolz, NVIDIA
- Neil Henning, AMD
- Tobias Hector, AMD
- Faith Ekstrand, Intel
- Baldur Karlsson, Valve
- Jan-Harald Fredriksen, Arm

## [](#_description)Description

This extension allows the application to query a 64-bit buffer device address value for a buffer, which can be used to access the buffer memory via the `PhysicalStorageBuffer` storage class in the [`GL_EXT_buffer_reference`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference.txt) GLSL extension and [`SPV_KHR_physical_storage_buffer`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_physical_storage_buffer.html) SPIR-V extension.

Another way to describe this extension is that it adds “pointers to buffer memory in shaders”. By calling [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html) with a `VkBuffer`, it will return a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value which represents the address of the start of the buffer.

[vkGetBufferOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferOpaqueCaptureAddress.html) and [vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html) allow opaque addresses for buffers and memory objects to be queried for the current process. A trace capture and replay tool can then supply these addresses to be used at replay time to match the addresses used when the trace was captured. To enable tools to insert these queries, new memory allocation flags must be specified for memory objects that will be bound to buffers accessed via the `PhysicalStorageBuffer` storage class. **Note that this mechanism is intended only to support capture/replay tools, and is not recommended for use in other applications.**

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. However, if Vulkan 1.2 is supported and this extension is not, the `bufferDeviceAddress` feature is optional. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

If Vulkan 1.3 is supported, support for the `bufferDeviceAddress` capability is required.

## [](#_new_commands)New Commands

- [vkGetBufferDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddressKHR.html)
- [vkGetBufferOpaqueCaptureAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferOpaqueCaptureAddressKHR.html)
- [vkGetDeviceMemoryOpaqueCaptureAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryOpaqueCaptureAddressKHR.html)

## [](#_new_structures)New Structures

- [VkBufferDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfoKHR.html)
- [VkDeviceMemoryOpaqueCaptureAddressInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOpaqueCaptureAddressInfoKHR.html)
- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html):
  
  - [VkBufferOpaqueCaptureAddressCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfoKHR.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkMemoryOpaqueCaptureAddressAllocateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryOpaqueCaptureAddressAllocateInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceBufferDeviceAddressFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBufferDeviceAddressFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME`
- `VK_KHR_BUFFER_DEVICE_ADDRESS_SPEC_VERSION`
- Extending [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html):
  
  - `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_KHR`
- Extending [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagBits.html):
  
  - `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT_KHR`
  - `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR`
  - `VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`PhysicalStorageBufferAddresses`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-PhysicalStorageBufferAddresses)

## [](#_examples)Examples

There are various use cases this extensions is designed for. It is required for ray tracing, useful for DX12 portability, and it allows storing buffer addresses in memory (enabling creating more complex data structures).

This extension can also be used to hardcode a dedicated debug channel into all shaders without impacting other descriptor limits by querying a buffer device address at startup and pushing that into shaders as a runtime constant (e.g. specialization constant).

There are examples of usage in the [`GL_EXT_buffer_reference`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference.txt) spec for how to use this in a high-level shading language such as GLSL. The [`GL_EXT_buffer_reference2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference2.txt) and [`GL_EXT_buffer_reference_uvec2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference_uvec2.txt) extensions were added to help cover a few edge cases missed by the original `GL_EXT_buffer_reference`.

## [](#_version_history)Version History

- Revision 1, 2019-06-24 (Jan-Harald Fredriksen)
  
  - Internal revisions based on VK\_EXT\_buffer\_device\_address

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_buffer_device_address).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0