# PFN_vkDebugReportCallbackEXT(3) Manual Page

## Name

PFN_vkDebugReportCallbackEXT - Application-defined debug report callback
function



## <a href="#_c_specification" class="anchor"></a>C Specification

The prototype for the
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)::`pfnCallback`
function implemented by the application is:

``` c
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

## <a href="#_parameters" class="anchor"></a>Parameters

- `flags` specifies the
  [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagBitsEXT.html) that
  triggered this callback.

- `objectType` is a
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html) value
  specifying the type of object being used or created at the time the
  event was triggered.

- `object` is the object where the issue was detected. If `objectType`
  is `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT`, `object` is undefined.

- `location` is a component (layer, driver, loader) defined value
  specifying the *location* of the trigger. This is an **optional**
  value.

- `messageCode` is a layer-defined value indicating what test triggered
  this callback.

- `pLayerPrefix` is a null-terminated UTF-8 string that is an
  abbreviation of the name of the component making the callback.
  `pLayerPrefix` is only valid for the duration of the callback.

- `pMessage` is a null-terminated UTF-8 string detailing the trigger
  conditions. `pMessage` is only valid for the duration of the callback.

- `pUserData` is the user data given when the
  [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html) was created.

## <a href="#_description" class="anchor"></a>Description

The callback **must** not call `vkDestroyDebugReportCallbackEXT`.

The callback returns a [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), which is interpreted
in a layer-specified manner. The application **should** always return
`VK_FALSE`. The `VK_TRUE` value is reserved for use in layer
development.

`object` **must** be a Vulkan object or
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html). If `objectType` is not
`VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT` and `object` is not
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `object` **must** be a Vulkan
object of the corresponding type associated with `objectType` as defined
in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debug-report-object-types"
class="bare" target="_blank"
rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debug-report-object-types</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html),
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html),
[VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagsEXT.html),
[VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PFN_vkDebugReportCallbackEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
