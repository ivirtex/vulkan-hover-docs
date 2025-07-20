# vkSetBufferCollectionImageConstraintsFUCHSIA(3) Manual Page

## Name

vkSetBufferCollectionImageConstraintsFUCHSIA - Set image-based constraints for a buffer collection



## [](#_c_specification)C Specification

Setting the constraints on the buffer collection initiates the format negotiation and allocation of the buffer collection. To set the constraints on a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) buffer collection, call:

```c++
// Provided by VK_FUCHSIA_buffer_collection
VkResult vkSetBufferCollectionImageConstraintsFUCHSIA(
    VkDevice                                    device,
    VkBufferCollectionFUCHSIA                   collection,
    const VkImageConstraintsInfoFUCHSIA*        pImageConstraintsInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device
- `collection` is the [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- `pImageConstraintsInfo` is a pointer to a [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html) structure

## [](#_description)Description

`vkSetBufferCollectionImageConstraintsFUCHSIA` **may** fail if `pImageConstraintsInfo->formatConstraintsCount` is larger than the implementation-defined limit. If that occurs, [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html) will return `VK_ERROR_INITIALIZATION_FAILED`.

`vkSetBufferCollectionImageConstraintsFUCHSIA` **may** fail if the implementation does not support any of the formats described by the `pImageConstraintsInfo` structure. If that occurs, [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html) will return `VK_ERROR_FORMAT_NOT_SUPPORTED`.

Valid Usage

- [](#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-06394)VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-06394  
  `vkSetBufferCollectionImageConstraintsFUCHSIA` or `vkSetBufferCollectionBufferConstraintsFUCHSIA` **must** not have already been called on `collection`

Valid Usage (Implicit)

- [](#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-device-parameter)VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parameter)VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parameter  
  `collection` **must** be a valid [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- [](#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-pImageConstraintsInfo-parameter)VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-pImageConstraintsInfo-parameter  
  `pImageConstraintsInfo` **must** be a valid pointer to a valid [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html) structure
- [](#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parent)VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parent  
  `collection` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_FORMAT_NOT_SUPPORTED`
- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetBufferCollectionImageConstraintsFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0