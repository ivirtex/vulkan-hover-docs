# vkDebugReportMessageEXT(3) Manual Page

## Name

vkDebugReportMessageEXT - Inject a message into a debug stream



## <a href="#_c_specification" class="anchor"></a>C Specification

To inject its own messages into the debug stream, call:

``` c
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

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the debug stream’s [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html).

- `flags` specifies the
  [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagBitsEXT.html)
  classification of this event/message.

- `objectType` is a
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html)
  specifying the type of object being used or created at the time the
  event was triggered.

- `object` is the object where the issue was detected. `object` **can**
  be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) if there is no object
  associated with the event.

- `location` is an application-defined value.

- `messageCode` is an application-defined value.

- `pLayerPrefix` is the abbreviation of the component making this
  event/message.

- `pMessage` is a null-terminated UTF-8 string detailing the trigger
  conditions.

## <a href="#_description" class="anchor"></a>Description

The call will propagate through the layers and generate callback(s) as
indicated by the message’s flags. The parameters are passed on to the
callback in addition to the `pUserData` value that was defined at the
time the callback was registered.

Valid Usage

- <a href="#VUID-vkDebugReportMessageEXT-object-01241"
  id="VUID-vkDebugReportMessageEXT-object-01241"></a>
  VUID-vkDebugReportMessageEXT-object-01241  
  `object` **must** be a Vulkan object or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkDebugReportMessageEXT-objectType-01498"
  id="VUID-vkDebugReportMessageEXT-objectType-01498"></a>
  VUID-vkDebugReportMessageEXT-objectType-01498  
  If `objectType` is not `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT` and
  `object` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `object`
  **must** be a Vulkan object of the corresponding type associated with
  `objectType` as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debug-report-object-types"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debug-report-object-types</a>

Valid Usage (Implicit)

- <a href="#VUID-vkDebugReportMessageEXT-instance-parameter"
  id="VUID-vkDebugReportMessageEXT-instance-parameter"></a>
  VUID-vkDebugReportMessageEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkDebugReportMessageEXT-flags-parameter"
  id="VUID-vkDebugReportMessageEXT-flags-parameter"></a>
  VUID-vkDebugReportMessageEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagBitsEXT.html) values

- <a href="#VUID-vkDebugReportMessageEXT-flags-requiredbitmask"
  id="VUID-vkDebugReportMessageEXT-flags-requiredbitmask"></a>
  VUID-vkDebugReportMessageEXT-flags-requiredbitmask  
  `flags` **must** not be `0`

- <a href="#VUID-vkDebugReportMessageEXT-objectType-parameter"
  id="VUID-vkDebugReportMessageEXT-objectType-parameter"></a>
  VUID-vkDebugReportMessageEXT-objectType-parameter  
  `objectType` **must** be a valid
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html) value

- <a href="#VUID-vkDebugReportMessageEXT-pLayerPrefix-parameter"
  id="VUID-vkDebugReportMessageEXT-pLayerPrefix-parameter"></a>
  VUID-vkDebugReportMessageEXT-pLayerPrefix-parameter  
  `pLayerPrefix` **must** be a null-terminated UTF-8 string

- <a href="#VUID-vkDebugReportMessageEXT-pMessage-parameter"
  id="VUID-vkDebugReportMessageEXT-pMessage-parameter"></a>
  VUID-vkDebugReportMessageEXT-pMessage-parameter  
  `pMessage` **must** be a null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html),
[VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagsEXT.html),
[VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDebugReportMessageEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
