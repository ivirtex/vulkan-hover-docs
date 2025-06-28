# PFN\_vkVoidFunction(3) Manual Page

## Name

PFN\_vkVoidFunction - Placeholder function pointer type returned by queries



## [](#_c_specification)C Specification

The definition of [PFN\_vkVoidFunction](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkVoidFunction.html) is:

```c++
// Provided by VK_VERSION_1_0
typedef void (VKAPI_PTR *PFN_vkVoidFunction)(void);
```

## [](#_description)Description

This type is returned from command function pointer queries, and **must** be cast to an actual command function pointer before use.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [vkGetDeviceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceProcAddr.html), [vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetInstanceProcAddr.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PFN_vkVoidFunction)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0