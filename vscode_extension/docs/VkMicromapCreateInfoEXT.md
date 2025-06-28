# VkMicromapCreateInfoEXT(3) Manual Page

## Name

VkMicromapCreateInfoEXT - Structure specifying the parameters of a newly created micromap object



## [](#_c_specification)C Specification

The `VkMicromapCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkMicromapCreateFlagsEXT    createFlags;
    VkBuffer                    buffer;
    VkDeviceSize                offset;
    VkDeviceSize                size;
    VkMicromapTypeEXT           type;
    VkDeviceAddress             deviceAddress;
} VkMicromapCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `createFlags` is a bitmask of [VkMicromapCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateFlagBitsEXT.html) specifying additional creation parameters of the micromap.
- `buffer` is the buffer on which the micromap will be stored.
- `offset` is an offset in bytes from the base address of the buffer at which the micromap will be stored, and **must** be a multiple of `256`.
- `size` is the size required for the micromap.
- `type` is a [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) value specifying the type of micromap that will be created.
- `deviceAddress` is the device address requested for the micromap if the [`micromapCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromapCaptureReplay) feature is being used.

## [](#_description)Description

If `deviceAddress` is zero, no specific address is requested.

If `deviceAddress` is not zero, `deviceAddress` **must** be an address retrieved from an identically created micromap on the same implementation. The micromap **must** also be placed on an identically created `buffer` and at the same `offset`.

Applications **should** avoid creating micromaps with application-provided addresses and implementation-provided addresses in the same process, to reduce the likelihood of `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR` errors.

Note

The expected usage for this is that a trace capture/replay tool will add the `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` flag to all buffers that use `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT`, and will add `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` to all buffers used as storage for a micromap where `deviceAddress` is not zero. This also means that the tool will need to add `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` to memory allocations to allow the flag to be set where the application may not have otherwise required it. During capture the tool will save the queried opaque device addresses in the trace. During replay, the buffers will be created specifying the original address so any address values stored in the trace data will remain valid.

Implementations are expected to separate such buffers in the GPU address space so normal allocations will avoid using these addresses. Applications and tools should avoid mixing application-provided and implementation-provided addresses for buffers created with `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`, to avoid address space allocation conflicts.

If the micromap will be the target of a build operation, the required size for a micromap **can** be queried with [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMicromapBuildSizesEXT.html).

Valid Usage

- [](#VUID-VkMicromapCreateInfoEXT-deviceAddress-07433)VUID-VkMicromapCreateInfoEXT-deviceAddress-07433  
  If `deviceAddress` is not zero, `createFlags` **must** include `VK_MICROMAP_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT`
- [](#VUID-VkMicromapCreateInfoEXT-createFlags-07434)VUID-VkMicromapCreateInfoEXT-createFlags-07434  
  If `createFlags` includes `VK_MICROMAP_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT`, [VkPhysicalDeviceOpacityMicromapFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpacityMicromapFeaturesEXT.html)::`micromapCaptureReplay` **must** be `VK_TRUE`
- [](#VUID-VkMicromapCreateInfoEXT-buffer-07435)VUID-VkMicromapCreateInfoEXT-buffer-07435  
  `buffer` **must** have been created with a `usage` value containing `VK_BUFFER_USAGE_MICROMAP_STORAGE_BIT_EXT`
- [](#VUID-VkMicromapCreateInfoEXT-buffer-07436)VUID-VkMicromapCreateInfoEXT-buffer-07436  
  `buffer` **must** not have been created with `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`
- [](#VUID-VkMicromapCreateInfoEXT-offset-07437)VUID-VkMicromapCreateInfoEXT-offset-07437  
  The sum of `offset` and `size` **must** be less than or equal to the size of `buffer`
- [](#VUID-VkMicromapCreateInfoEXT-offset-07438)VUID-VkMicromapCreateInfoEXT-offset-07438  
  `offset` **must** be a multiple of `256` bytes

Valid Usage (Implicit)

- [](#VUID-VkMicromapCreateInfoEXT-sType-sType)VUID-VkMicromapCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MICROMAP_CREATE_INFO_EXT`
- [](#VUID-VkMicromapCreateInfoEXT-pNext-pNext)VUID-VkMicromapCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMicromapCreateInfoEXT-createFlags-parameter)VUID-VkMicromapCreateInfoEXT-createFlags-parameter  
  `createFlags` **must** be a valid combination of [VkMicromapCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateFlagBitsEXT.html) values
- [](#VUID-VkMicromapCreateInfoEXT-buffer-parameter)VUID-VkMicromapCreateInfoEXT-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkMicromapCreateInfoEXT-type-parameter)VUID-VkMicromapCreateInfoEXT-type-parameter  
  `type` **must** be a valid [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkMicromapCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateFlagsEXT.html), [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMicromapEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMicromapCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0