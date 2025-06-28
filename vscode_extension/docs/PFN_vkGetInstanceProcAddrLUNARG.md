# PFN\_vkGetInstanceProcAddrLUNARG(3) Manual Page

## Name

PFN\_vkGetInstanceProcAddrLUNARG - Type definition for vkGetInstanceProcAddr



## [](#_c_specification)C Specification

The type of [PFN\_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkGetInstanceProcAddrLUNARG.html) is:

```c++
// Provided by VK_LUNARG_direct_driver_loading
typedef PFN_vkVoidFunction (VKAPI_PTR *PFN_vkGetInstanceProcAddrLUNARG)(
    VkInstance instance, const char* pName);
```

## [](#_parameters)Parameters

- `instance` is a [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle.
- `pName` is the name of a Vulkan command.

## [](#_description)Description

This type is compatible with the type of a pointer to the [vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetInstanceProcAddr.html) command, but is used only to specify device driver addresses in [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingInfoLUNARG.html)::`pfnGetInstanceProcAddr`.

Note

This type exists only because of limitations in the XML schema and processing scripts, and its name may change in the future. Ideally we would use the `PFN_vkGetInstanceProcAddr` type generated in the `vulkan_core.h` header.

## [](#_see_also)See Also

[VK\_LUNARG\_direct\_driver\_loading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_LUNARG_direct_driver_loading.html), [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingInfoLUNARG.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PFN_vkGetInstanceProcAddrLUNARG)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0