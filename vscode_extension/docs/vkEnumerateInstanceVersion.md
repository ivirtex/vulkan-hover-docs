# vkEnumerateInstanceVersion(3) Manual Page

## Name

vkEnumerateInstanceVersion - Query instance-level version before instance creation



## [](#_c_specification)C Specification

To query the version of instance-level functionality supported by the implementation, call:

```c++
// Provided by VK_VERSION_1_1
VkResult vkEnumerateInstanceVersion(
    uint32_t*                                   pApiVersion);
```

## [](#_parameters)Parameters

- `pApiVersion` is a pointer to a `uint32_t`, which is the version of Vulkan supported by instance-level functionality, encoded as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-coreversions-versionnumbers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-coreversions-versionnumbers).

## [](#_description)Description

Note

The intended behavior of [vkEnumerateInstanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumerateInstanceVersion.html) is that an implementation **should** not need to perform memory allocations and **should** unconditionally return `VK_SUCCESS`. The loader, and any enabled layers, **may** return `VK_ERROR_OUT_OF_HOST_MEMORY` in the case of a failed memory allocation.

Valid Usage (Implicit)

- [](#VUID-vkEnumerateInstanceVersion-pApiVersion-parameter)VUID-vkEnumerateInstanceVersion-pApiVersion-parameter  
  `pApiVersion` **must** be a valid pointer to a `uint32_t` value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkEnumerateInstanceVersion)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0