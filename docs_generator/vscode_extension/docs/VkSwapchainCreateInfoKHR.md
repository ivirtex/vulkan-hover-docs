# VkSwapchainCreateInfoKHR(3) Manual Page

## Name

VkSwapchainCreateInfoKHR - Structure specifying parameters of a newly created swapchain object



## [](#_c_specification)C Specification

The `VkSwapchainCreateInfoKHR` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagBitsKHR.html) indicating parameters of the swapchain creation.
- `surface` is the surface onto which the swapchain will present images. If the creation succeeds, the swapchain becomes associated with `surface`.
- `minImageCount` is the minimum number of presentable images that the application needs. The implementation will either create the swapchain with at least that many images, or it will fail to create the swapchain.
- `imageFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value specifying the format the swapchain image(s) will be created with.
- `imageColorSpace` is a [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html) value specifying the way the swapchain interprets image data.
- `imageExtent` is the size (in pixels) of the swapchain image(s). The behavior is platform-dependent if the image extent does not match the surface’s `currentExtent` as returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR`.
  
  Note
  
  On some platforms, it is normal that `maxImageExtent` **may** become `(0, 0)`, for example when the window is minimized. In such a case, it is not possible to create a swapchain due to the Valid Usage requirements , unless scaling is selected through [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoEXT.html), if supported .
- `imageArrayLayers` is the number of views in a multiview/stereo surface. For non-stereoscopic-3D applications, this value is 1.
- `imageUsage` is a bitmask of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) describing the intended usage of the (acquired) swapchain images.
- `imageSharingMode` is the sharing mode used for the image(s) of the swapchain.
- `queueFamilyIndexCount` is the number of queue families having access to the image(s) of the swapchain when `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`.
- `pQueueFamilyIndices` is a pointer to an array of queue family indices having access to the images(s) of the swapchain when `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`.
- `preTransform` is a [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html) value describing the transform, relative to the presentation engine’s natural orientation, applied to the image content prior to presentation. If it does not match the `currentTransform` value returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR`, the presentation engine will transform the image content as part of the presentation operation.
- `compositeAlpha` is a [VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompositeAlphaFlagBitsKHR.html) value indicating the alpha compositing mode to use when this surface is composited together with other surfaces on certain window systems.
- `presentMode` is the presentation mode the swapchain will use. A swapchain’s present mode determines how incoming present requests will be processed and queued internally.
- `clipped` specifies whether the Vulkan implementation is allowed to discard rendering operations that affect regions of the surface that are not visible.
  
  - If `clipped` is `VK_TRUE`, the presentable images associated with the swapchain **may** not own all of their pixels. Pixels in the presentable images that correspond to regions of the target surface obscured by another window on the desktop, or subject to some other clipping mechanism will have undefined content when read back. Fragment shaders **may** not execute for these pixels, and thus any side effects they would have had will not occur. Setting `VK_TRUE` does not guarantee any clipping will occur, but allows more efficient presentation methods to be used on some platforms.
  - If `clipped` is `VK_FALSE`, presentable images associated with the swapchain will own all of the pixels they contain.
    
    Note
    
    Applications **should** set this value to `VK_TRUE` if they do not expect to read back the content of presentable images before presenting them or after reacquiring them, and if their fragment shaders do not have any side effects that require them to run for all pixels in the presentable image.
- `oldSwapchain` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), or the existing non-retired swapchain currently associated with `surface`. Providing a valid `oldSwapchain` **may** aid in the resource reuse, and also allows the application to still present any images that are already acquired from it.

## [](#_description)Description

Upon calling `vkCreateSwapchainKHR` with an `oldSwapchain` that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `oldSwapchain` is retired — even if creation of the new swapchain fails. The new swapchain is created in the non-retired state whether or not `oldSwapchain` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html).

Upon calling `vkCreateSwapchainKHR` with an `oldSwapchain` that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), any images from `oldSwapchain` that are not acquired by the application **may** be freed by the implementation, which **may** occur even if creation of the new swapchain fails. The application **can** destroy `oldSwapchain` to free all memory associated with `oldSwapchain`.

Note

Multiple retired swapchains **can** be associated with the same `VkSurfaceKHR` through multiple uses of `oldSwapchain` that outnumber calls to [vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySwapchainKHR.html).

After `oldSwapchain` is retired, the application **can** pass to [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) any images it had already acquired from `oldSwapchain`. E.g., an application may present an image from the old swapchain before an image from the new swapchain is ready to be presented. As usual, [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) **may** fail if `oldSwapchain` has entered a state that causes `VK_ERROR_OUT_OF_DATE_KHR` to be returned.

The application **can** continue to use a shared presentable image obtained from `oldSwapchain` until a presentable image is acquired from the new swapchain, as long as it has not entered a state that causes it to return `VK_ERROR_OUT_OF_DATE_KHR`.

Valid Usage

- [](#VUID-VkSwapchainCreateInfoKHR-surface-01270)VUID-VkSwapchainCreateInfoKHR-surface-01270  
  `surface` **must** be a surface that is supported by the device as determined using [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
- [](#VUID-VkSwapchainCreateInfoKHR-minImageCount-01272)VUID-VkSwapchainCreateInfoKHR-minImageCount-01272  
  `minImageCount` **must** be less than or equal to the value returned in the `maxImageCount` member of the `VkSurfaceCapabilitiesKHR` structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for the surface if the returned `maxImageCount` is not zero
- [](#VUID-VkSwapchainCreateInfoKHR-swapchainMaintenance1-10155)VUID-VkSwapchainCreateInfoKHR-swapchainMaintenance1-10155  
  If the [`swapchainMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-swapchainMaintenance1) feature is not enabled, then the `pNext` chain **must** not include a [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html) structure
- [](#VUID-VkSwapchainCreateInfoKHR-presentMode-02839)VUID-VkSwapchainCreateInfoKHR-presentMode-02839  
  If `presentMode` is not `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` nor `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`, then `minImageCount` **must** be greater than or equal to the value returned in the `minImageCount` member of the `VkSurfaceCapabilitiesKHR` structure returned by [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html) for the surface
- [](#VUID-VkSwapchainCreateInfoKHR-minImageCount-01383)VUID-VkSwapchainCreateInfoKHR-minImageCount-01383  
  `minImageCount` **must** be `1` if `presentMode` is either `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`
- [](#VUID-VkSwapchainCreateInfoKHR-imageFormat-01273)VUID-VkSwapchainCreateInfoKHR-imageFormat-01273  
  `imageFormat` and `imageColorSpace` **must** match the `format` and `colorSpace` members, respectively, of one of the `VkSurfaceFormatKHR` structures returned by `vkGetPhysicalDeviceSurfaceFormatsKHR` for the surface
- [](#VUID-VkSwapchainCreateInfoKHR-pNext-07781)VUID-VkSwapchainCreateInfoKHR-pNext-07781  
  If a [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoEXT.html) structure was not included in the `pNext` chain, or it is included and [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)::`scalingBehavior` is zero then `imageExtent` **must** be between `minImageExtent` and `maxImageExtent`, inclusive, where `minImageExtent` and `maxImageExtent` are members of the `VkSurfaceCapabilitiesKHR` structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for the surface
- [](#VUID-VkSwapchainCreateInfoKHR-pNext-07782)VUID-VkSwapchainCreateInfoKHR-pNext-07782  
  If a [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoEXT.html) structure was included in the `pNext` chain and [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)::`scalingBehavior` is not zero then `imageExtent` **must** be between `minScaledImageExtent` and `maxScaledImageExtent`, inclusive, where `minScaledImageExtent` and `maxScaledImageExtent` are members of the `VkSurfacePresentScalingCapabilitiesEXT` structure returned by `vkGetPhysicalDeviceSurfaceCapabilities2KHR` for the surface and `presentMode`
- [](#VUID-VkSwapchainCreateInfoKHR-swapchainMaintenance1-10157)VUID-VkSwapchainCreateInfoKHR-swapchainMaintenance1-10157  
  If the [`swapchainMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-swapchainMaintenance1) feature is not enabled, then `flags` **must** not include `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT`
- [](#VUID-VkSwapchainCreateInfoKHR-imageExtent-01689)VUID-VkSwapchainCreateInfoKHR-imageExtent-01689  
  `imageExtent` members `width` and `height` **must** both be non-zero
- [](#VUID-VkSwapchainCreateInfoKHR-imageArrayLayers-01275)VUID-VkSwapchainCreateInfoKHR-imageArrayLayers-01275  
  `imageArrayLayers` **must** be greater than `0` and less than or equal to the `maxImageArrayLayers` member of the `VkSurfaceCapabilitiesKHR` structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for the surface
- [](#VUID-VkSwapchainCreateInfoKHR-presentMode-01427)VUID-VkSwapchainCreateInfoKHR-presentMode-01427  
  If `presentMode` is `VK_PRESENT_MODE_FIFO_LATEST_READY_EXT`, `VK_PRESENT_MODE_IMMEDIATE_KHR`, `VK_PRESENT_MODE_MAILBOX_KHR`, `VK_PRESENT_MODE_FIFO_KHR` or `VK_PRESENT_MODE_FIFO_RELAXED_KHR`, `imageUsage` **must** be a subset of the supported usage flags present in the `supportedUsageFlags` member of the [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html) structure returned by [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html) for `surface`
- [](#VUID-VkSwapchainCreateInfoKHR-imageUsage-01384)VUID-VkSwapchainCreateInfoKHR-imageUsage-01384  
  If `presentMode` is `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`, `imageUsage` **must** be a subset of the supported usage flags present in the `sharedPresentSupportedUsageFlags` member of the [VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html) structure returned by [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html) for `surface`
- [](#VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01277)VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01277  
  If `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`, `pQueueFamilyIndices` **must** be a valid pointer to an array of `queueFamilyIndexCount` `uint32_t` values
- [](#VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01278)VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01278  
  If `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`, `queueFamilyIndexCount` **must** be greater than `1`
- [](#VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01428)VUID-VkSwapchainCreateInfoKHR-imageSharingMode-01428  
  If `imageSharingMode` is `VK_SHARING_MODE_CONCURRENT`, each element of `pQueueFamilyIndices` **must** be unique and **must** be less than `pQueueFamilyPropertyCount` returned by either [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html) or [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html) for the `physicalDevice` that was used to create `device`
- [](#VUID-VkSwapchainCreateInfoKHR-preTransform-01279)VUID-VkSwapchainCreateInfoKHR-preTransform-01279  
  `preTransform` **must** be one of the bits present in the `supportedTransforms` member of the `VkSurfaceCapabilitiesKHR` structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for the surface
- [](#VUID-VkSwapchainCreateInfoKHR-compositeAlpha-01280)VUID-VkSwapchainCreateInfoKHR-compositeAlpha-01280  
  `compositeAlpha` **must** be one of the bits present in the `supportedCompositeAlpha` member of the `VkSurfaceCapabilitiesKHR` structure returned by `vkGetPhysicalDeviceSurfaceCapabilitiesKHR` for the surface
- [](#VUID-VkSwapchainCreateInfoKHR-presentMode-01281)VUID-VkSwapchainCreateInfoKHR-presentMode-01281  
  `presentMode` **must** be one of the [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values returned by `vkGetPhysicalDeviceSurfacePresentModesKHR` for the surface
- [](#VUID-VkSwapchainCreateInfoKHR-presentModeFifoLatestReady-10161)VUID-VkSwapchainCreateInfoKHR-presentModeFifoLatestReady-10161  
  If the [`presentModeFifoLatestReady`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentModeFifoLatestReady) feature is not enabled, `presentMode` **must** not be `VK_PRESENT_MODE_FIFO_LATEST_READY_EXT`
- [](#VUID-VkSwapchainCreateInfoKHR-physicalDeviceCount-01429)VUID-VkSwapchainCreateInfoKHR-physicalDeviceCount-01429  
  If the logical device was created with [VkDeviceGroupDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupDeviceCreateInfo.html)::`physicalDeviceCount` equal to 1, `flags` **must** not contain `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR`
- [](#VUID-VkSwapchainCreateInfoKHR-oldSwapchain-01933)VUID-VkSwapchainCreateInfoKHR-oldSwapchain-01933  
  If `oldSwapchain` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `oldSwapchain` **must** be a non-retired swapchain associated with native window referred to by `surface`
- [](#VUID-VkSwapchainCreateInfoKHR-imageFormat-01778)VUID-VkSwapchainCreateInfoKHR-imageFormat-01778  
  The [implied image creation parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#swapchain-wsi-image-create-info) of the swapchain **must** be supported as reported by [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties.html)
- [](#VUID-VkSwapchainCreateInfoKHR-flags-03168)VUID-VkSwapchainCreateInfoKHR-flags-03168  
  If `flags` contains `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` then the `pNext` chain **must** include a [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html) structure with a `viewFormatCount` greater than zero and `pViewFormats` **must** have an element equal to `imageFormat`
- [](#VUID-VkSwapchainCreateInfoKHR-pNext-04099)VUID-VkSwapchainCreateInfoKHR-pNext-04099  
  If a [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html) structure was included in the `pNext` chain and [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount` is not zero then all of the formats in [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats` **must** be compatible with the `format` as described in the [compatibility table](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility)
- [](#VUID-VkSwapchainCreateInfoKHR-flags-04100)VUID-VkSwapchainCreateInfoKHR-flags-04100  
  If `flags` does not contain `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` and the `pNext` chain include a [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html) structure then [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount` **must** be `0` or `1`
- [](#VUID-VkSwapchainCreateInfoKHR-flags-03187)VUID-VkSwapchainCreateInfoKHR-flags-03187  
  If `flags` contains `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR`, then `VkSurfaceProtectedCapabilitiesKHR`::`supportsProtected` **must** be `VK_TRUE` in the [VkSurfaceProtectedCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceProtectedCapabilitiesKHR.html) structure returned by [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html) for `surface`
- [](#VUID-VkSwapchainCreateInfoKHR-pNext-02679)VUID-VkSwapchainCreateInfoKHR-pNext-02679  
  If the `pNext` chain includes a [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html) structure with its `fullScreenExclusive` member set to `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`, and `surface` was created using [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWin32SurfaceKHR.html), a [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html) structure **must** be included in the `pNext` chain
- [](#VUID-VkSwapchainCreateInfoKHR-pNext-06752)VUID-VkSwapchainCreateInfoKHR-pNext-06752  
  If the [`imageCompressionControlSwapchain`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-imageCompressionControlSwapchain) feature is not enabled, the `pNext` chain **must** not include an [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html) structure

Valid Usage (Implicit)

- [](#VUID-VkSwapchainCreateInfoKHR-sType-sType)VUID-VkSwapchainCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR`
- [](#VUID-VkSwapchainCreateInfoKHR-pNext-pNext)VUID-VkSwapchainCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDeviceGroupSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html), [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html), [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html), [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html), [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html), [VkSwapchainCounterCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCounterCreateInfoEXT.html), [VkSwapchainDisplayNativeHdrCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainDisplayNativeHdrCreateInfoAMD.html), [VkSwapchainLatencyCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainLatencyCreateInfoNV.html), [VkSwapchainPresentBarrierCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentBarrierCreateInfoNV.html), [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html), or [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)
- [](#VUID-VkSwapchainCreateInfoKHR-sType-unique)VUID-VkSwapchainCreateInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkSwapchainCreateInfoKHR-flags-parameter)VUID-VkSwapchainCreateInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagBitsKHR.html) values
- [](#VUID-VkSwapchainCreateInfoKHR-surface-parameter)VUID-VkSwapchainCreateInfoKHR-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-VkSwapchainCreateInfoKHR-imageFormat-parameter)VUID-VkSwapchainCreateInfoKHR-imageFormat-parameter  
  `imageFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkSwapchainCreateInfoKHR-imageColorSpace-parameter)VUID-VkSwapchainCreateInfoKHR-imageColorSpace-parameter  
  `imageColorSpace` **must** be a valid [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html) value
- [](#VUID-VkSwapchainCreateInfoKHR-imageUsage-parameter)VUID-VkSwapchainCreateInfoKHR-imageUsage-parameter  
  `imageUsage` **must** be a valid combination of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) values
- [](#VUID-VkSwapchainCreateInfoKHR-imageUsage-requiredbitmask)VUID-VkSwapchainCreateInfoKHR-imageUsage-requiredbitmask  
  `imageUsage` **must** not be `0`
- [](#VUID-VkSwapchainCreateInfoKHR-imageSharingMode-parameter)VUID-VkSwapchainCreateInfoKHR-imageSharingMode-parameter  
  `imageSharingMode` **must** be a valid [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html) value
- [](#VUID-VkSwapchainCreateInfoKHR-preTransform-parameter)VUID-VkSwapchainCreateInfoKHR-preTransform-parameter  
  `preTransform` **must** be a valid [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html) value
- [](#VUID-VkSwapchainCreateInfoKHR-compositeAlpha-parameter)VUID-VkSwapchainCreateInfoKHR-compositeAlpha-parameter  
  `compositeAlpha` **must** be a valid [VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompositeAlphaFlagBitsKHR.html) value
- [](#VUID-VkSwapchainCreateInfoKHR-presentMode-parameter)VUID-VkSwapchainCreateInfoKHR-presentMode-parameter  
  `presentMode` **must** be a valid [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) value
- [](#VUID-VkSwapchainCreateInfoKHR-oldSwapchain-parameter)VUID-VkSwapchainCreateInfoKHR-oldSwapchain-parameter  
  If `oldSwapchain` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `oldSwapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-VkSwapchainCreateInfoKHR-commonparent)VUID-VkSwapchainCreateInfoKHR-commonparent  
  Both of `oldSwapchain`, and `surface` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

Host Synchronization

- Host access to `surface` **must** be externally synchronized
- Host access to `oldSwapchain` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html), [VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompositeAlphaFlagBitsKHR.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html), [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html), [VkSwapchainCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagsKHR.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html), [vkCreateSharedSwapchainsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSharedSwapchainsKHR.html), [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0