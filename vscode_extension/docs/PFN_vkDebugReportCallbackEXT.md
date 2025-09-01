# PFN\_vkDebugReportCallbackEXT(3) Manual Page

## Name

PFN\_vkDebugReportCallbackEXT - Application-defined debug report callback function



## [](#_c_specification)C Specification

The prototype for the [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html)::`pfnCallback` function implemented by the application is:

```c++
// Provided by VK_EXT_debug_report
typedef VkBool32 (VKAPI_PTR *PFN_vkDebugReportCallbackEXT)(
    VkDebugReportFlagsEXT                       flags,
    VkDebugReportObjectTypeEXT                  objectType,
    uint64_t                                    object,
    size_t                                      location,
    int32_t                                     messageCode,
    const char*                                 pLayerPrefix,
    const char*                                 pMessage,
    void*                                       pUserData);
```

## [](#_parameters)Parameters

- `flags` specifies the [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagBitsEXT.html) that triggered this callback.
- `objectType` is a [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) value specifying the type of object being used or created at the time the event was triggered.
- `object` is the object where the issue was detected. If `objectType` is `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT`, `object` is undefined.
- `location` is a component (layer, driver, loader) defined value specifying the *location* of the trigger. This is an **optional** value.
- `messageCode` is a layer-defined value indicating what test triggered this callback.
- `pLayerPrefix` is a null-terminated UTF-8 string that is an abbreviation of the name of the component making the callback. `pLayerPrefix` is only valid for the duration of the callback.
- `pMessage` is a null-terminated UTF-8 string detailing the trigger conditions. `pMessage` is only valid for the duration of the callback.
- `pUserData` is the user data given when the [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html) was created.

## [](#_description)Description

The callback **must** not call `vkDestroyDebugReportCallbackEXT`.

The callback returns a [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), which is interpreted in a layer-specified manner. The application **should** always return `VK_FALSE`. The `VK_TRUE` value is reserved for use in layer development.

`object` **must** be a Vulkan object or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html). If `objectType` is not `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT` and `object` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `object` **must** be a Vulkan object of the corresponding type associated with `objectType` as defined in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debug-report-object-types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debug-report-object-types).

## [](#_see_also)See Also

[VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html), [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html), [VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagsEXT.html), [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PFN_vkDebugReportCallbackEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0