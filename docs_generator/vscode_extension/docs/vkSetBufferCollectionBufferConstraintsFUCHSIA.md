# vkSetBufferCollectionBufferConstraintsFUCHSIA(3) Manual Page

## Name

vkSetBufferCollectionBufferConstraintsFUCHSIA - Set buffer-based constraints for a buffer collection



## [](#_c_specification)C Specification

To set the constraints on a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) buffer collection, call:

```c++
// Provided by VK_FUCHSIA_buffer_collection
VkResult vkSetBufferCollectionBufferConstraintsFUCHSIA(
    VkDevice                                    device,
    VkBufferCollectionFUCHSIA                   collection,
    const VkBufferConstraintsInfoFUCHSIA*       pBufferConstraintsInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device
- `collection` is the [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- `pBufferConstraintsInfo` is a pointer to a [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferConstraintsInfoFUCHSIA.html) structure

## [](#_description)Description

`vkSetBufferCollectionBufferConstraintsFUCHSIA` **may** fail if the implementation does not support the constraints specified in the `bufferCollectionConstraints` structure. If that occurs, [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html) will return `VK_ERROR_FORMAT_NOT_SUPPORTED`.

Valid Usage

- [](#VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-collection-06403)VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-collection-06403  
  `vkSetBufferCollectionImageConstraintsFUCHSIA` or `vkSetBufferCollectionBufferConstraintsFUCHSIA` **must** not have already been called on `collection`

Valid Usage (Implicit)

- [](#VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-device-parameter)VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-collection-parameter)VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-collection-parameter  
  `collection` **must** be a valid [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) handle
- [](#VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-pBufferConstraintsInfo-parameter)VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-pBufferConstraintsInfo-parameter  
  `pBufferConstraintsInfo` **must** be a valid pointer to a valid [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferConstraintsInfoFUCHSIA.html) structure
- [](#VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-collection-parent)VUID-vkSetBufferCollectionBufferConstraintsFUCHSIA-collection-parent  
  `collection` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_FORMAT_NOT_SUPPORTED`

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html), [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferConstraintsInfoFUCHSIA.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetBufferCollectionBufferConstraintsFUCHSIA)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0