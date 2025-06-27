# VK_KHR_present_wait(3) Manual Page

## Name

VK_KHR_present_wait - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

249

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  
and  
[VK_KHR_present_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_present_id.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Keith Packard <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_present_wait%5D%20@keithp%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_present_wait%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>keithp</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-05-15

**IP Status**  
No known IP claims.

**Contributors**  
- Keith Packard, Valve

- Ian Elliott, Google

- Tobias Hector, AMD

- Daniel Stone, Collabora

## <a href="#_description" class="anchor"></a>Description

This device extension allows an application that uses the
[`VK_KHR_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) extension to wait for
present operations to complete. An application can use this to monitor
and control the pacing of the application by managing the number of
outstanding images yet to be presented.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkWaitForPresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitForPresentKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePresentWaitFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePresentWaitFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_PRESENT_WAIT_EXTENSION_NAME`

- `VK_KHR_PRESENT_WAIT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_WAIT_FEATURES_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) When does the wait finish?

**RESOLVED**. The wait will finish when the present is visible to the
user. There is no requirement for any precise timing relationship
between the presentation of the image to the user, but implementations
**should** signal the wait as close as possible to the presentation of
the first pixel in the new image to the user.

2\) Should this use fences or other existing synchronization mechanism.

**RESOLVED**. Because display and rendering are often implemented in
separate drivers, this extension will provide a separate synchronization
API.

3\) Should this extension share present identification with other
extensions?

**RESOLVED**. Yes. A new extension, VK_KHR_present_id, should be created
to provide a shared structure for presentation identifiers.

4\) What happens when presentations complete out of order wrt calls to
vkQueuePresent? This could happen if the semaphores for the
presentations were ready out of order.

**OPTION A**: Require that when a PresentId is set that the driver
ensure that images are always presented in the order of calls to
vkQueuePresent.

**OPTION B**: Finish both waits when the earliest present completes.
This will complete the later present wait earlier than the actual
presentation. This should be the easiest to implement as the driver need
only track the largest present ID completed. This is also the 'natural'
consequence of interpreting the existing vkWaitForPresentKHR
specificationn.

**OPTION C**: Finish both waits when both have completed. This will
complete the earlier presentation later than the actual presentation
time. This is allowed by the current specification as there is no
precise timing requirement for when the presentId value is updated. This
requires slightly more complexity in the driver as it will need to track
all outstanding presentId values.

## <a href="#_examples" class="anchor"></a>Examples

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-02-19 (Keith Packard)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_present_wait"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
