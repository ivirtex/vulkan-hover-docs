# VK_KHR_buffer_device_address(3) Manual Page

## Name

VK_KHR_buffer_device_address - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

258

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_physical_storage_buffer](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_physical_storage_buffer.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_buffer_device_address%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_buffer_device_address%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-06-24

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_EXT_buffer_reference`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference.txt)
  and
  [`GL_EXT_buffer_reference2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference2.txt)
  and
  [`GL_EXT_buffer_reference_uvec2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference_uvec2.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

- Neil Henning, AMD

- Tobias Hector, AMD

- Faith Ekstrand, Intel

- Baldur Karlsson, Valve

- Jan-Harald Fredriksen, Arm

## <a href="#_description" class="anchor"></a>Description

This extension allows the application to query a 64-bit buffer device
address value for a buffer, which can be used to access the buffer
memory via the `PhysicalStorageBuffer` storage class in the
[`GL_EXT_buffer_reference`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference.txt)
GLSL extension and
[`SPV_KHR_physical_storage_buffer`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_physical_storage_buffer.html)
SPIR-V extension.

Another way to describe this extension is that it adds “pointers to
buffer memory in shaders”. By calling
[vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html) with a
`VkBuffer`, it will return a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html)
value which represents the address of the start of the buffer.

[vkGetBufferOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html)
and
[vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html)
allow opaque addresses for buffers and memory objects to be queried for
the current process. A trace capture and replay tool can then supply
these addresses to be used at replay time to match the addresses used
when the trace was captured. To enable tools to insert these queries,
new memory allocation flags must be specified for memory objects that
will be bound to buffers accessed via the `PhysicalStorageBuffer`
storage class. **Note that this mechanism is intended only to support
capture/replay tools, and is not recommended for use in other
applications.**

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. However, if Vulkan 1.2 is supported and this
extension is not, the `bufferDeviceAddress` feature is optional. The
original type, enum and command names are still available as aliases of
the core functionality.

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Support for the `bufferDeviceAddress` feature is mandatory in Vulkan
1.3, regardless of whether this extension is supported.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetBufferDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddressKHR.html)

- [vkGetBufferOpaqueCaptureAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddressKHR.html)

- [vkGetDeviceMemoryOpaqueCaptureAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddressKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBufferDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfoKHR.html)

- [VkDeviceMemoryOpaqueCaptureAddressInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOpaqueCaptureAddressInfoKHR.html)

- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html):

  - [VkBufferOpaqueCaptureAddressCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfoKHR.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkMemoryOpaqueCaptureAddressAllocateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfoKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceBufferDeviceAddressFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME`

- `VK_KHR_BUFFER_DEVICE_ADDRESS_SPEC_VERSION`

- Extending [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html):

  - `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR`

- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html):

  - `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_KHR`

- Extending [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html):

  - `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT_KHR`

  - `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR`

  - `VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-PhysicalStorageBufferAddresses"
  target="_blank"
  rel="noopener"><code>PhysicalStorageBufferAddresses</code></a>

## <a href="#_examples" class="anchor"></a>Examples

There are various use cases this extensions is designed for. It is
required for ray tracing, useful for DX12 portability, and it allows
storing buffer addresses in memory (enabling creating more complex data
structures).

This extension can also be used to hardcode a dedicated debug channel
into all shaders without impacting other descriptor limits by querying a
buffer device address at startup and pushing that into shaders as a
runtime constant (e.g. specialization constant).

There are examples of usage in the
[`GL_EXT_buffer_reference`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference.txt)
spec for how to use this in a high-level shading language such as GLSL.
The
[`GL_EXT_buffer_reference2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference2.txt)
and
[`GL_EXT_buffer_reference_uvec2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_buffer_reference_uvec2.txt)
extensions were added to help cover a few edge cases missed by the
original `GL_EXT_buffer_reference`.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-06-24 (Jan-Harald Fredriksen)

  - Internal revisions based on VK_EXT_buffer_device_address

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_buffer_device_address"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
