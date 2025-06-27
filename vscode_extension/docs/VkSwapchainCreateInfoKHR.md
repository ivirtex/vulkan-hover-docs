# VkSwapchainCreateInfoKHR(3) Manual Page

## Name

VkSwapchainCreateInfoKHR - Structure specifying parameters of a newly
created swapchain object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSwapchainCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_swapchain
typedef struct VkSwapchainCreateInfoKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkSwapchainCreateFlagsKHR        flags;
    VkSurfaceKHR                     surface;
    uint32_t                         minImageCount;
    VkFormat                         imageFormat;
    VkColorSpaceKHR                  imageColorSpace;
    VkExtent2D                       imageExtent;
    uint32_t                         imageArrayLayers;
    VkImageUsageFlags                imageUsage;
    VkSharingMode                    imageSharingMode;
    uint32_t                         queueFamilyIndexCount;
    const uint32_t*                  pQueueFamilyIndices;
    VkSurfaceTransformFlagBitsKHR    preTransform;
    VkCompositeAlphaFlagBitsKHR      compositeAlpha;
    VkPresentModeKHR                 presentMode;
    VkBool32                         clipped;
    VkSwapchainKHR                   oldSwapchain;
} VkSwapchainCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagBitsKHR.html)
  indicating parameters of the swapchain creation.

- `surface` is the surface onto which the swapchain will present images.
  If the creation succeeds, the swapchain becomes associated with
  `surface`.

- `minImageCount` is the minimum number of presentable images that the
  application needs. The implementation will either create the swapchain
  with at least that many images, or it will fail to create the
  swapchain.

- `imageFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value specifying the
  format the swapchain image(s) will be created with.

- `imageColorSpace` is a [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html) value
  specifying the way the swapchain interprets image data.

- `imageExtent` is the size (in pixels) of the swapchain image(s). The
  behavior is platform-dependent if the image extent does not match the
  surface’s `currentExtent` as returned by
  `vkGetPhysicalDeviceSurfaceCapabilitiesKHR`.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>On some platforms, it is normal that <code>maxImageExtent</code>
  <strong>may</strong> become <code>(0, 0)</code>, for example when the
  window is minimized. In such a case, it is not possible to create a
  swapchain due to the Valid Usage requirements , unless scaling is
  selected through <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html">VkSwapchainPresentScalingCreateInfoEXT</a>,
  if supported .</p></td>
  </tr>
  </tbody>
  </table>

- `imageArrayLayers` is the number of views in a multiview/stereo
  surface. For non-stereoscopic-3D applications, this value is 1.

- `imageUsage` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) describing the
  intended usage of the (acquired) swapchain images.

- `imageSharingMode` is the sharing mode used for the image(s) of the
  swapchain.

- `queueFamilyIndexCount` is the number of queue families having access
  to the image(s) of the swapchain when `imageSharingMode` is
  `VK_SHARING_MODE_CONCURRENT`.

- `pQueueFamilyIndices` is a pointer to an array of queue family indices
  having access to the images(s) of the swapchain when
  `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`.

- `preTransform` is a
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  value describing the transform, relative to the presentation engine’s
  natural orientation, applied to the image content prior to
  presentation. If it does not match the `currentTransform` value
  returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR`, the
  presentation engine will transform the image content as part of the
  presentation operation.

- `compositeAlpha` is a
  [VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html) value
  indicating the alpha compositing mode to use when this surface is
  composited together with other surfaces on certain window systems.

- `presentMode` is the presentation mode the swapchain will use. A
  swapchain’s present mode determines how incoming present requests will
  be processed and queued internally.

- `clipped` specifies whether the Vulkan implementation is allowed to
  discard rendering operations that affect regions of the surface that
  are not visible.

  - If set to `VK_TRUE`, the presentable images associated with the
    swapchain **may** not own all of their pixels. Pixels in the
    presentable images that correspond to regions of the target surface
    obscured by another window on the desktop, or subject to some other
    clipping mechanism will have undefined content when read back.
    Fragment shaders **may** not execute for these pixels, and thus any
    side effects they would have had will not occur. Setting `VK_TRUE`
    does not guarantee any clipping will occur, but allows more
    efficient presentation methods to be used on some platforms.

  - If set to `VK_FALSE`, presentable images associated with the
    swapchain will own all of the pixels they contain.

    <table>
    <colgroup>
    <col style="width: 50%" />
    <col style="width: 50%" />
    </colgroup>
    <tbody>
    <tr>
    <td class="icon"><em></em></td>
    <td class="content">Note
    <p>Applications <strong>should</strong> set this value to
    <code>VK_TRUE</code> if they do not expect to read back the content of
    presentable images before presenting them or after reacquiring them, and
    if their fragment shaders do not have any side effects that require them
    to run for all pixels in the presentable image.</p></td>
    </tr>
    </tbody>
    </table>

- `oldSwapchain` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), or the
  existing non-retired swapchain currently associated with `surface`.
  Providing a valid `oldSwapchain` **may** aid in the resource reuse,
  and also allows the application to still present any images that are
  already acquired from it.

## <a href="#_description" class="anchor"></a>Description

Upon calling `vkCreateSwapchainKHR` with an `oldSwapchain` that is not
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `oldSwapchain` is retired — even
if creation of the new swapchain fails. The new swapchain is created in
the non-retired state whether or not `oldSwapchain` is
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html).

Upon calling `vkCreateSwapchainKHR` with an `oldSwapchain` that is not
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), any images from `oldSwapchain`
that are not acquired by the application **may** be freed by the
implementation, which **may** occur even if creation of the new
swapchain fails. The application **can** destroy `oldSwapchain` to free
all memory associated with `oldSwapchain`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Multiple retired swapchains <strong>can</strong> be associated with
the same <code>VkSurfaceKHR</code> through multiple uses of
<code>oldSwapchain</code> that outnumber calls to <a
href="vkDestroySwapchainKHR.html">vkDestroySwapchainKHR</a>.</p>
<p>After <code>oldSwapchain</code> is retired, the application
<strong>can</strong> pass to <a
href="vkQueuePresentKHR.html">vkQueuePresentKHR</a> any images it had
already acquired from <code>oldSwapchain</code>. E.g., an application
may present an image from the old swapchain before an image from the new
swapchain is ready to be presented. As usual, <a
href="vkQueuePresentKHR.html">vkQueuePresentKHR</a> <strong>may</strong>
fail if <code>oldSwapchain</code> has entered a state that causes
<code>VK_ERROR_OUT_OF_DATE_KHR</code> to be returned.</p>
<p>The application <strong>can</strong> continue to use a shared
presentable image obtained from <code>oldSwapchain</code> until a
presentable image is acquired from the new swapchain, as long as it has
not entered a state that causes it to return
<code>VK_ERROR_OUT_OF_DATE_KHR</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkSwapchainCreateInfoKHR-surface-01270"
  id="VUID-VkSwapchainCreateInfoKHR-surface-01270"></a>
  VUID-VkSwapchainCreateInfoKHR-surface-01270  
  `surface` **must** be a surface that is supported by the device as
  determined using
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)

- <a href="#VUID-VkSwapchainCreateInfoKHR-minImageCount-01272"
  id="VUID-VkSwapchainCreateInfoKHR-minImageCount-01272"></a>
  VUID-VkSwapchainCreateInfoKHR-minImageCount-01272  
  `minImageCount` **must** be less than or equal to the value returned
  in the `maxImageCount` member of the `VkSurfaceCapabilitiesKHR`
  structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for
  the surface if the returned `maxImageCount` is not zero

- <a href="#VUID-VkSwapchainCreateInfoKHR-presentMode-02839"
  id="VUID-VkSwapchainCreateInfoKHR-presentMode-02839"></a>
  VUID-VkSwapchainCreateInfoKHR-presentMode-02839  
  If `presentMode` is not `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`
  nor `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`, then
  `minImageCount` **must** be greater than or equal to the value
  returned in the `minImageCount` member of the
  `VkSurfaceCapabilitiesKHR` structure returned by
  [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html)
  for the surface

- <a href="#VUID-VkSwapchainCreateInfoKHR-minImageCount-01383"
  id="VUID-VkSwapchainCreateInfoKHR-minImageCount-01383"></a>
  VUID-VkSwapchainCreateInfoKHR-minImageCount-01383  
  `minImageCount` **must** be `1` if `presentMode` is either
  `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or
  `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageFormat-01273"
  id="VUID-VkSwapchainCreateInfoKHR-imageFormat-01273"></a>
  VUID-VkSwapchainCreateInfoKHR-imageFormat-01273  
  `imageFormat` and `imageColorSpace` **must** match the `format` and
  `colorSpace` members, respectively, of one of the `VkSurfaceFormatKHR`
  structures returned by `vkGetPhysicalDeviceSurfaceFormatsKHR` for the
  surface

- <a href="#VUID-VkSwapchainCreateInfoKHR-pNext-07781"
  id="VUID-VkSwapchainCreateInfoKHR-pNext-07781"></a>
  VUID-VkSwapchainCreateInfoKHR-pNext-07781  
  If a
  [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)
  structure was not included in the `pNext` chain, or it is included and
  [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)::`scalingBehavior`
  is zero then `imageExtent` **must** be between `minImageExtent` and
  `maxImageExtent`, inclusive, where `minImageExtent` and
  `maxImageExtent` are members of the `VkSurfaceCapabilitiesKHR`
  structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for
  the surface

- <a href="#VUID-VkSwapchainCreateInfoKHR-pNext-07782"
  id="VUID-VkSwapchainCreateInfoKHR-pNext-07782"></a>
  VUID-VkSwapchainCreateInfoKHR-pNext-07782  
  If a
  [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)
  structure was included in the `pNext` chain and
  [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)::`scalingBehavior`
  is not zero then `imageExtent` **must** be between
  `minScaledImageExtent` and `maxScaledImageExtent`, inclusive, where
  `minScaledImageExtent` and `maxScaledImageExtent` are members of the
  `VkSurfacePresentScalingCapabilitiesEXT` structure returned by
  `vkGetPhysicalDeviceSurfaceCapabilities2KHR` for the surface and
  `presentMode`

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageExtent-01689"
  id="VUID-VkSwapchainCreateInfoKHR-imageExtent-01689"></a>
  VUID-VkSwapchainCreateInfoKHR-imageExtent-01689  
  `imageExtent` members `width` and `height` **must** both be non-zero

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageArrayLayers-01275"
  id="VUID-VkSwapchainCreateInfoKHR-imageArrayLayers-01275"></a>
  VUID-VkSwapchainCreateInfoKHR-imageArrayLayers-01275  
  `imageArrayLayers` **must** be greater than `0` and less than or equal
  to the `maxImageArrayLayers` member of the `VkSurfaceCapabilitiesKHR`
  structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for
  the surface

- <a href="#VUID-VkSwapchainCreateInfoKHR-presentMode-01427"
  id="VUID-VkSwapchainCreateInfoKHR-presentMode-01427"></a>
  VUID-VkSwapchainCreateInfoKHR-presentMode-01427  
  If `presentMode` is `VK_PRESENT_MODE_IMMEDIATE_KHR`,
  `VK_PRESENT_MODE_MAILBOX_KHR`, `VK_PRESENT_MODE_FIFO_KHR` or
  `VK_PRESENT_MODE_FIFO_RELAXED_KHR`, `imageUsage` **must** be a subset
  of the supported usage flags present in the `supportedUsageFlags`
  member of the
  [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html) structure
  returned by
  [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html)
  for `surface`

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageUsage-01384"
  id="VUID-VkSwapchainCreateInfoKHR-imageUsage-01384"></a>
  VUID-VkSwapchainCreateInfoKHR-imageUsage-01384  
  If `presentMode` is `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or
  `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`, `imageUsage` **must**
  be a subset of the supported usage flags present in the
  `sharedPresentSupportedUsageFlags` member of the
  [VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html)
  structure returned by
  [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html)
  for `surface`

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01277"
  id="VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01277"></a>
  VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01277  
  If `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`,
  `pQueueFamilyIndices` **must** be a valid pointer to an array of
  `queueFamilyIndexCount` `uint32_t` values

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01278"
  id="VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01278"></a>
  VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01278  
  If `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`,
  `queueFamilyIndexCount` **must** be greater than `1`

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01428"
  id="VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01428"></a>
  VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01428  
  If `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`, each element of
  `pQueueFamilyIndices` **must** be unique and **must** be less than
  `pQueueFamilyPropertyCount` returned by either
  [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html)
  or
  [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html)
  for the `physicalDevice` that was used to create `device`

- <a href="#VUID-VkSwapchainCreateInfoKHR-preTransform-01279"
  id="VUID-VkSwapchainCreateInfoKHR-preTransform-01279"></a>
  VUID-VkSwapchainCreateInfoKHR-preTransform-01279  
  `preTransform` **must** be one of the bits present in the
  `supportedTransforms` member of the `VkSurfaceCapabilitiesKHR`
  structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for
  the surface

- <a href="#VUID-VkSwapchainCreateInfoKHR-compositeAlpha-01280"
  id="VUID-VkSwapchainCreateInfoKHR-compositeAlpha-01280"></a>
  VUID-VkSwapchainCreateInfoKHR-compositeAlpha-01280  
  `compositeAlpha` **must** be one of the bits present in the
  `supportedCompositeAlpha` member of the `VkSurfaceCapabilitiesKHR`
  structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for
  the surface

- <a href="#VUID-VkSwapchainCreateInfoKHR-presentMode-01281"
  id="VUID-VkSwapchainCreateInfoKHR-presentMode-01281"></a>
  VUID-VkSwapchainCreateInfoKHR-presentMode-01281  
  `presentMode` **must** be one of the
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) values returned by
  `vkGetPhysicalDeviceSurfacePresentModesKHR` for the surface

- <a href="#VUID-VkSwapchainCreateInfoKHR-physicalDeviceCount-01429"
  id="VUID-VkSwapchainCreateInfoKHR-physicalDeviceCount-01429"></a>
  VUID-VkSwapchainCreateInfoKHR-physicalDeviceCount-01429  
  If the logical device was created with
  [VkDeviceGroupDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html)::`physicalDeviceCount`
  equal to 1, `flags` **must** not contain
  `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR`

- <a href="#VUID-VkSwapchainCreateInfoKHR-oldSwapchain-01933"
  id="VUID-VkSwapchainCreateInfoKHR-oldSwapchain-01933"></a>
  VUID-VkSwapchainCreateInfoKHR-oldSwapchain-01933  
  If `oldSwapchain` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `oldSwapchain` **must** be a non-retired swapchain associated with
  native window referred to by `surface`

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageFormat-01778"
  id="VUID-VkSwapchainCreateInfoKHR-imageFormat-01778"></a>
  VUID-VkSwapchainCreateInfoKHR-imageFormat-01778  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#swapchain-wsi-image-create-info"
  target="_blank" rel="noopener">implied image creation parameters</a>
  of the swapchain **must** be supported as reported by
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)

- <a href="#VUID-VkSwapchainCreateInfoKHR-flags-03168"
  id="VUID-VkSwapchainCreateInfoKHR-flags-03168"></a>
  VUID-VkSwapchainCreateInfoKHR-flags-03168  
  If `flags` contains `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` then
  the `pNext` chain **must** include a
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure with a `viewFormatCount` greater than zero and
  `pViewFormats` **must** have an element equal to `imageFormat`

- <a href="#VUID-VkSwapchainCreateInfoKHR-pNext-04099"
  id="VUID-VkSwapchainCreateInfoKHR-pNext-04099"></a>
  VUID-VkSwapchainCreateInfoKHR-pNext-04099  
  If a [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure was included in the `pNext` chain and
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount`
  is not zero then all of the formats in
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats`
  **must** be compatible with the `format` as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility"
  target="_blank" rel="noopener">compatibility table</a>

- <a href="#VUID-VkSwapchainCreateInfoKHR-flags-04100"
  id="VUID-VkSwapchainCreateInfoKHR-flags-04100"></a>
  VUID-VkSwapchainCreateInfoKHR-flags-04100  
  If `flags` does not contain
  `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` and the `pNext` chain
  include a
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure then
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount`
  **must** be `0` or `1`

- <a href="#VUID-VkSwapchainCreateInfoKHR-flags-03187"
  id="VUID-VkSwapchainCreateInfoKHR-flags-03187"></a>
  VUID-VkSwapchainCreateInfoKHR-flags-03187  
  If `flags` contains `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR`, then
  `VkSurfaceProtectedCapabilitiesKHR`::`supportsProtected` **must** be
  `VK_TRUE` in the
  [VkSurfaceProtectedCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceProtectedCapabilitiesKHR.html)
  structure returned by
  [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html)
  for `surface`

- <a href="#VUID-VkSwapchainCreateInfoKHR-pNext-02679"
  id="VUID-VkSwapchainCreateInfoKHR-pNext-02679"></a>
  VUID-VkSwapchainCreateInfoKHR-pNext-02679  
  If the `pNext` chain includes a
  [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)
  structure with its `fullScreenExclusive` member set to
  `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`, and `surface`
  was created using
  [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWin32SurfaceKHR.html), a
  [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html)
  structure **must** be included in the `pNext` chain

- <a href="#VUID-VkSwapchainCreateInfoKHR-pNext-06752"
  id="VUID-VkSwapchainCreateInfoKHR-pNext-06752"></a>
  VUID-VkSwapchainCreateInfoKHR-pNext-06752  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-imageCompressionControlSwapchain"
  target="_blank"
  rel="noopener"><code>imageCompressionControlSwapchain</code></a>
  feature is not enabled, the `pNext` chain **must** not include an
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainCreateInfoKHR-sType-sType"
  id="VUID-VkSwapchainCreateInfoKHR-sType-sType"></a>
  VUID-VkSwapchainCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR`

- <a href="#VUID-VkSwapchainCreateInfoKHR-pNext-pNext"
  id="VUID-VkSwapchainCreateInfoKHR-pNext-pNext"></a>
  VUID-VkSwapchainCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDeviceGroupSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html),
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html),
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html),
  [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html),
  [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html),
  [VkSwapchainCounterCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCounterCreateInfoEXT.html),
  [VkSwapchainDisplayNativeHdrCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainDisplayNativeHdrCreateInfoAMD.html),
  [VkSwapchainLatencyCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainLatencyCreateInfoNV.html),
  [VkSwapchainPresentBarrierCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentBarrierCreateInfoNV.html),
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html),
  or
  [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)

- <a href="#VUID-VkSwapchainCreateInfoKHR-sType-unique"
  id="VUID-VkSwapchainCreateInfoKHR-sType-unique"></a>
  VUID-VkSwapchainCreateInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkSwapchainCreateInfoKHR-flags-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-flags-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagBitsKHR.html)
  values

- <a href="#VUID-VkSwapchainCreateInfoKHR-surface-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-surface-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageFormat-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-imageFormat-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-imageFormat-parameter  
  `imageFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageColorSpace-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-imageColorSpace-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-imageColorSpace-parameter  
  `imageColorSpace` **must** be a valid
  [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html) value

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageUsage-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-imageUsage-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-imageUsage-parameter  
  `imageUsage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageUsage-requiredbitmask"
  id="VUID-VkSwapchainCreateInfoKHR-imageUsage-requiredbitmask"></a>
  VUID-VkSwapchainCreateInfoKHR-imageUsage-requiredbitmask  
  `imageUsage` **must** not be `0`

- <a href="#VUID-VkSwapchainCreateInfoKHR-imageSharingMode-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-imageSharingMode-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-imageSharingMode-parameter  
  `imageSharingMode` **must** be a valid
  [VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html) value

- <a href="#VUID-VkSwapchainCreateInfoKHR-preTransform-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-preTransform-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-preTransform-parameter  
  `preTransform` **must** be a valid
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  value

- <a href="#VUID-VkSwapchainCreateInfoKHR-compositeAlpha-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-compositeAlpha-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-compositeAlpha-parameter  
  `compositeAlpha` **must** be a valid
  [VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html) value

- <a href="#VUID-VkSwapchainCreateInfoKHR-presentMode-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-presentMode-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-presentMode-parameter  
  `presentMode` **must** be a valid
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) value

- <a href="#VUID-VkSwapchainCreateInfoKHR-oldSwapchain-parameter"
  id="VUID-VkSwapchainCreateInfoKHR-oldSwapchain-parameter"></a>
  VUID-VkSwapchainCreateInfoKHR-oldSwapchain-parameter  
  If `oldSwapchain` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `oldSwapchain` **must** be a valid
  [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) handle

- <a href="#VUID-VkSwapchainCreateInfoKHR-commonparent"
  id="VUID-VkSwapchainCreateInfoKHR-commonparent"></a>
  VUID-VkSwapchainCreateInfoKHR-commonparent  
  Both of `oldSwapchain`, and `surface` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html),
[VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html),
[VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html),
[VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html),
[VkSwapchainCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagsKHR.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html),
[vkCreateSharedSwapchainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSharedSwapchainsKHR.html),
[vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
