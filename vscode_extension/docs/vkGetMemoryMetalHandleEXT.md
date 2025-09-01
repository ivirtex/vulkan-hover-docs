# vkGetMemoryMetalHandleEXT(3) Manual Page

## Name

vkGetMemoryMetalHandleEXT - Get a Metal handle for a memory object



## [](#_c_specification)C Specification

To export a Metal handle representing the payload of a Vulkan device memory object, call:

```c++
// Provided by VK_EXT_external_memory_metal
VkResult vkGetMemoryMetalHandleEXT(
    VkDevice                                    device,
    const VkMemoryGetMetalHandleInfoEXT*        pGetMetalHandleInfo,
    void**                                      pHandle);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the device memory being exported.
- `pGetMetalHandleInfo` is a pointer to a [VkMemoryGetMetalHandleInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetMetalHandleInfoEXT.html) structure containing parameters of the export operation.
- `pHandle` will return the Metal handle representing the payload of the device memory object.

## [](#_description)Description

Unless the app retains the handle object returned by the call, the lifespan will be the same as the associated `VkDeviceMemory`.

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryMetalHandleEXT-device-parameter)VUID-vkGetMemoryMetalHandleEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryMetalHandleEXT-pGetMetalHandleInfo-parameter)VUID-vkGetMemoryMetalHandleEXT-pGetMetalHandleInfo-parameter  
  `pGetMetalHandleInfo` **must** be a valid pointer to a valid [VkMemoryGetMetalHandleInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetMetalHandleInfoEXT.html) structure
- [](#VUID-vkGetMemoryMetalHandleEXT-pHandle-parameter)VUID-vkGetMemoryMetalHandleEXT-pHandle-parameter  
  `pHandle` **must** be a valid pointer to a pointer value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_external\_memory\_metal](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_external_memory_metal.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryGetMetalHandleInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetMetalHandleInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryMetalHandleEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0