# VkDebugUtilsMessageSeverityFlagBitsEXT(3) Manual Page

## Name

VkDebugUtilsMessageSeverityFlagBitsEXT - Bitmask specifying which
severities of events cause a debug messenger callback



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)::`messageSeverity`,
specifying event severities which cause a debug messenger to call the
callback, are:

``` c
// Provided by VK_EXT_debug_utils
typedef enum VkDebugUtilsMessageSeverityFlagBitsEXT {
    VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT = 0x00000001,
    VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT = 0x00000010,
    VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT = 0x00000100,
    VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT = 0x00001000,
} VkDebugUtilsMessageSeverityFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT` specifies the most
  verbose output indicating all diagnostic messages from the Vulkan
  loader, layers, and drivers should be captured.

- `VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT` specifies an
  informational message such as resource details that may be handy when
  debugging an application.

- `VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT` specifies use of
  Vulkan that **may** expose an app bug. Such cases may not be
  immediately harmful, such as a fragment shader outputting to a
  location with no attachment. Other cases **may** point to behavior
  that is almost certainly bad when unintended such as using an image
  whose memory has not been filled. In general if you see a warning but
  you know that the behavior is intended/desired, then simply ignore the
  warning.

- `VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT` specifies that the
  application has violated a valid usage condition of the specification.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The values of <a
href="VkDebugUtilsMessageSeverityFlagBitsEXT.html">VkDebugUtilsMessageSeverityFlagBitsEXT</a>
are sorted based on severity. The higher the flag value, the more severe
the message. This allows for simple boolean operation comparisons when
looking at <a
href="VkDebugUtilsMessageSeverityFlagBitsEXT.html">VkDebugUtilsMessageSeverityFlagBitsEXT</a>
values.</p>
<p>For example:</p>
<pre class="c"><code>    if (messageSeverity &gt;= VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT) {
        // Do something for warnings and errors
    }</code></pre>
<p>In addition, space has been left between the enums to allow for later
addition of new severities in between the existing values.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsMessageSeverityFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagsEXT.html),
[vkSubmitDebugUtilsMessageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSubmitDebugUtilsMessageEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugUtilsMessageSeverityFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
