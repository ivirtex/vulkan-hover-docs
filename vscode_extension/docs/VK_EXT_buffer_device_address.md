# VK\_EXT\_buffer\_device\_address(3) Manual Page

## Name

VK\_EXT\_buffer\_device\_address - device extension



## [](#_registered_extension_number)Registered Extension Number

245

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_physical\_storage\_buffer](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_physical_storage_buffer.html)

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_buffer_device_address%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_buffer_device_address%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-01-06

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_EXT_buffer_reference`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference.txt) and [`GLSL_EXT_buffer_reference_uvec2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference_uvec2.txt)

**Contributors**

- Jeff Bolz, NVIDIA
- Neil Henning, AMD
- Tobias Hector, AMD
- Faith Ekstrand, Intel
- Baldur Karlsson, Valve

## [](#_description)Description

This extension allows the application to query a 64-bit buffer device address value for a buffer, which can be used to access the buffer memory via the `PhysicalStorageBufferEXT` storage class in the [`GL_EXT_buffer_reference`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference.txt) GLSL extension and [`SPV_EXT_physical_storage_buffer`](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_physical_storage_buffer.html) SPIR-V extension.

It also allows buffer device addresses to be provided by a trace replay tool, so that it matches the address used when the trace was captured.

## [](#_new_commands)New Commands

- [vkGetBufferDeviceAddressEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddressEXT.html)

## [](#_new_structures)New Structures

- [VkBufferDeviceAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfoEXT.html)
- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html):
  
  - [VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceBufferAddressFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBufferAddressFeaturesEXT.html)
  - [VkPhysicalDeviceBufferDeviceAddressFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBufferDeviceAddressFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME`
- `VK_EXT_BUFFER_DEVICE_ADDRESS_SPEC_VERSION`
- Extending [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html):
  
  - `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_EXT`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_INVALID_DEVICE_ADDRESS_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_ADDRESS_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`PhysicalStorageBufferAddressesEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-PhysicalStorageBufferAddresses)

## [](#_issues)Issues

1\) Where is VK\_STRUCTURE\_TYPE\_PHYSICAL\_DEVICE\_BUFFER\_ADDRESS\_FEATURES\_EXT and VkPhysicalDeviceBufferAddressFeaturesEXT?

**RESOLVED**: They were renamed as `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT` and [VkPhysicalDeviceBufferDeviceAddressFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBufferDeviceAddressFeaturesEXT.html) accordingly for consistency. Even though, the old names can still be found in the generated header files for compatibility.

## [](#_version_history)Version History

- Revision 1, 2018-11-01 (Jeff Bolz)
  
  - Internal revisions
- Revision 2, 2019-01-06 (Jon Leech)
  
  - Minor updates to appendix for publication

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_buffer_device_address)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0