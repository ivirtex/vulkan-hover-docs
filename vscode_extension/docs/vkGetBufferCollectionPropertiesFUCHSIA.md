# vkGetBufferCollectionPropertiesFUCHSIA(3) Manual Page

## Name

vkGetBufferCollectionPropertiesFUCHSIA - Retrieve properties from a buffer collection



## [](#_c_specification)C Specification

After constraints have been set on the buffer collection by calling [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html) or [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html), call `vkGetBufferCollectionPropertiesFUCHSIA` to retrieve the negotiated and finalized properties of the buffer collection.

The call to `vkGetBufferCollectionPropertiesFUCHSIA` is synchronous. It waits for the Sysmem format negotiation and buffer collection allocation to complete before returning.

```c++
// Provided by VK_FUCHSIA_buffer_collection
VkResult vkGetBufferCollectionPropertiesFUCHSIA(
    VkDevice                                    device,
    VkBufferCollectionFUCHSIA                   collection,
    VkBufferCollectionPropertiesFUCHSIA*        pProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device handle
- `collection` is the [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- `pProperties` is a pointer to the retrieved [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html) struct

## [](#_description)Description

For image-based buffer collections, upon calling `vkGetBufferCollectionPropertiesFUCHSIA`, Sysmem will choose an element of the [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html)::`pImageCreateInfos` established by the preceding call to [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html). The index of the element chosen is stored in and can be retrieved from [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html)::`createInfoIndex`.

For buffer-based buffer collections, a single [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) is specified as [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferConstraintsInfoFUCHSIA.html)::`createInfo`. [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html)::`createInfoIndex` will therefore always be zero.

`vkGetBufferCollectionPropertiesFUCHSIA` **may** fail if Sysmem is unable to resolve the constraints of all of the participants in the buffer collection. If that occurs, `vkGetBufferCollectionPropertiesFUCHSIA` will return `VK_ERROR_INITIALIZATION_FAILED`.

Valid Usage

- [](#VUID-vkGetBufferCollectionPropertiesFUCHSIA-None-06405)VUID-vkGetBufferCollectionPropertiesFUCHSIA-None-06405  
  Prior to calling [vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html), the constraints on the buffer collection **must** have been set by either [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html) or [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html)

Valid Usage (Implicit)

- [](#VUID-vkGetBufferCollectionPropertiesFUCHSIA-device-parameter)VUID-vkGetBufferCollectionPropertiesFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parameter)VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parameter  
  `collection` **must** be a valid [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- [](#VUID-vkGetBufferCollectionPropertiesFUCHSIA-pProperties-parameter)VUID-vkGetBufferCollectionPropertiesFUCHSIA-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html) structure
- [](#VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parent)VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parent  
  `collection` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html), [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetBufferCollectionPropertiesFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0