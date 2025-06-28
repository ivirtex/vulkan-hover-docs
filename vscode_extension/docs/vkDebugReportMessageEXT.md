# vkDebugReportMessageEXT(3) Manual Page

## Name

vkDebugReportMessageEXT - Inject a message into a debug stream



## [](#_c_specification)C Specification

To inject its own messages into the debug stream, call:

```c++
// Provided by VK_EXT_debug_report
void vkDebugReportMessageEXT(
    VkInstance                                  instance,
    VkDebugReportFlagsEXT                       flags,
    VkDebugReportObjectTypeEXT                  objectType,
    uint64_t                                    object,
    size_t                                      location,
    int32_t                                     messageCode,
    const char*                                 pLayerPrefix,
    const char*                                 pMessage);
```

## [](#_parameters)Parameters

- `instance` is the debug stream’s [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html).
- `flags` specifies the [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagBitsEXT.html) classification of this event/message.
- `objectType` is a [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) specifying the type of object being used or created at the time the event was triggered.
- `object` is the object where the issue was detected. `object` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) if there is no object associated with the event.
- `location` is an application-defined value.
- `messageCode` is an application-defined value.
- `pLayerPrefix` is the abbreviation of the component making this event/message.
- `pMessage` is a null-terminated UTF-8 string detailing the trigger conditions.

## [](#_description)Description

The call will propagate through the layers and generate callback(s) as indicated by the message’s flags. The parameters are passed on to the callback in addition to the `pUserData` value that was defined at the time the callback was registered.

Valid Usage

- [](#VUID-vkDebugReportMessageEXT-object-01241)VUID-vkDebugReportMessageEXT-object-01241  
  `object` **must** be a Vulkan object or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkDebugReportMessageEXT-objectType-01498)VUID-vkDebugReportMessageEXT-objectType-01498  
  If `objectType` is not `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT` and `object` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `object` **must** be a Vulkan object of the corresponding type associated with `objectType` as defined in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debug-report-object-types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debug-report-object-types)

Valid Usage (Implicit)

- [](#VUID-vkDebugReportMessageEXT-instance-parameter)VUID-vkDebugReportMessageEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkDebugReportMessageEXT-flags-parameter)VUID-vkDebugReportMessageEXT-flags-parameter  
  `flags` **must** be a valid combination of [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagBitsEXT.html) values
- [](#VUID-vkDebugReportMessageEXT-flags-requiredbitmask)VUID-vkDebugReportMessageEXT-flags-requiredbitmask  
  `flags` **must** not be `0`
- [](#VUID-vkDebugReportMessageEXT-objectType-parameter)VUID-vkDebugReportMessageEXT-objectType-parameter  
  `objectType` **must** be a valid [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) value
- [](#VUID-vkDebugReportMessageEXT-pLayerPrefix-parameter)VUID-vkDebugReportMessageEXT-pLayerPrefix-parameter  
  `pLayerPrefix` **must** be a null-terminated UTF-8 string
- [](#VUID-vkDebugReportMessageEXT-pMessage-parameter)VUID-vkDebugReportMessageEXT-pMessage-parameter  
  `pMessage` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html), [VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagsEXT.html), [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDebugReportMessageEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0