# vkCreateMicromapEXT(3) Manual Page

## Name

vkCreateMicromapEXT - Create a new micromap object



## [](#_c_specification)C Specification

To create a micromap, call:

```c++
// Provided by VK_EXT_opacity_micromap
VkResult vkCreateMicromapEXT(
    VkDevice                                    device,
    const VkMicromapCreateInfoEXT*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkMicromapEXT*                              pMicromap);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the micromap object.
- `pCreateInfo` is a pointer to a [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html) structure containing parameters affecting creation of the micromap.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pMicromap` is a pointer to a [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle in which the resulting micromap object is returned.

## [](#_description)Description

Similar to other objects in Vulkan, the micromap creation merely creates an object with a specific “shape”. The type and quantity of geometry that can be built into a micromap is determined by the parameters of [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html).

The micromap data is stored in the object referred to by `VkMicromapCreateInfoEXT`::`buffer`. Once memory has been bound to that buffer, it **must** be populated by micromap build or micromap copy commands such as [vkCmdBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildMicromapsEXT.html), [vkBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBuildMicromapsEXT.html), [vkCmdCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapEXT.html), and [vkCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMicromapEXT.html).

Note

The expected usage for a trace capture/replay tool is that it will serialize and later deserialize the micromap data using micromap copy commands. During capture the tool will use [vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMicromapToMemoryEXT.html) or [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html) with a `mode` of `VK_COPY_MICROMAP_MODE_SERIALIZE_EXT`, and [vkCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToMicromapEXT.html) or [vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToMicromapEXT.html) with a `mode` of `VK_COPY_MICROMAP_MODE_DESERIALIZE_EXT` during replay.

The input buffers passed to micromap build commands will be referenced by the implementation for the duration of the command. Micromaps **must** be fully self-contained. The application **can** reuse or free any memory which was used by the command as an input or as scratch without affecting the results of a subsequent acceleration structure build using the micromap or traversal of that acceleration structure.

Valid Usage

- [](#VUID-vkCreateMicromapEXT-micromap-07430)VUID-vkCreateMicromapEXT-micromap-07430  
  The [`micromap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromap) feature **must** be enabled
- [](#VUID-vkCreateMicromapEXT-deviceAddress-07431)VUID-vkCreateMicromapEXT-deviceAddress-07431  
  If [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html)::`deviceAddress` is not zero, the [`micromapCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromapCaptureReplay) feature **must** be enabled
- [](#VUID-vkCreateMicromapEXT-device-07432)VUID-vkCreateMicromapEXT-device-07432  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCreateMicromapEXT-device-parameter)VUID-vkCreateMicromapEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateMicromapEXT-pCreateInfo-parameter)VUID-vkCreateMicromapEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html) structure
- [](#VUID-vkCreateMicromapEXT-pAllocator-parameter)VUID-vkCreateMicromapEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateMicromapEXT-pMicromap-parameter)VUID-vkCreateMicromapEXT-pMicromap-parameter  
  `pMicromap` **must** be a valid pointer to a [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle
- [](#VUID-vkCreateMicromapEXT-device-queuecount)VUID-vkCreateMicromapEXT-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateMicromapEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0